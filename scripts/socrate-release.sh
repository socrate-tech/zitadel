#!/usr/bin/env bash
# scripts/socrate-release.sh
#
# All-in-one release helper for the Socrate-patched Zitadel image.
#
# Wraps the workspace reset, install, nx pack, and tag/push steps into a
# single command. Auto-discovers the next patch revision from the target
# registry when --version is not supplied. Optionally switches to (or
# bootstraps) the fork patch branch corresponding to an upstream tag.
#
# Scope is deliberately limited: this script does NOT touch socrate-flux
# or socrate-k8s. Environment promotion stays a separate, reviewable step.
#
# Usage:
#   scripts/socrate-release.sh                              # auto-bumped Scaleway release on current branch
#   scripts/socrate-release.sh --target local               # smoke-test build
#   scripts/socrate-release.sh --version v4.7.6-4           # explicit version
#   scripts/socrate-release.sh --branch v4.7.6              # switch to fork branch v4.7.6-x (must match an upstream tag)
#   scripts/socrate-release.sh --branch v4.15.0             # bootstrap fork branch v4.15.0-x from upstream tag v4.15.0
#   scripts/socrate-release.sh --dry-run                    # rehearsal, no side effects

set -euo pipefail

trap 'rc=$?; (( rc != 0 )) && printf "\n\033[31mxx socrate-release failed (line %s, exit %s)\033[0m\n" "${LINENO}" "$rc" >&2' EXIT

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

SCALEWAY_REGISTRY="rg.fr-par.scw.cloud"
SCALEWAY_NAMESPACES=("cr-registry-dev" "cr-registry-qa" "cr-registry-staging" "cr-registry-prod")
SCALEWAY_PROBE_NAMESPACE="cr-registry-prod"
LOCAL_REGISTRY="registry.internal.local.socrate.ninja:5001"

TARGET="scaleway"
VERSION_OVERRIDE=""
UPSTREAM_TAG=""
DRY_RUN=false
SKIP_RESET=false
YES=false

usage() {
  sed -n '2,20p' "${BASH_SOURCE[0]}" | sed 's/^# \{0,1\}//'
  exit "${1:-0}"
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --target)     TARGET="${2:?missing value for --target}"; shift 2 ;;
    --version)    VERSION_OVERRIDE="${2:?missing value for --version}"; shift 2 ;;
    --branch)     UPSTREAM_TAG="${2:?missing value for --branch}"; shift 2 ;;
    --dry-run)    DRY_RUN=true; shift ;;
    --skip-reset) SKIP_RESET=true; shift ;;
    --yes|-y)     YES=true; shift ;;
    --help|-h)    usage 0 ;;
    *)            printf "Unknown argument: %s\n\n" "$1" >&2; usage 64 ;;
  esac
done

case "$TARGET" in
  local|scaleway) ;;
  *) printf "Invalid --target '%s' (expected: local|scaleway)\n" "$TARGET" >&2; exit 64 ;;
esac

if [[ -n "$UPSTREAM_TAG" && -n "$VERSION_OVERRIDE" ]]; then
  if [[ "$VERSION_OVERRIDE" =~ ^v([0-9]+\.[0-9]+\.[0-9]+)-[0-9]+$ ]]; then
    _version_base="${BASH_REMATCH[1]}"
    _upstream_base="${UPSTREAM_TAG#v}"
    if [[ "$_version_base" != "$_upstream_base" ]]; then
      printf "Inconsistent flags: --branch %s implies upstream base %s, but --version %s has base %s\n" \
        "$UPSTREAM_TAG" "$_upstream_base" "$VERSION_OVERRIDE" "$_version_base" >&2
      exit 64
    fi
  fi
fi

say()  { printf "\033[36m▶\033[0m %s\n" "$*"; }
warn() { printf "\033[33m! %s\033[0m\n" "$*" >&2; }
fail() { printf "\033[31mxx %s\033[0m\n" "$*" >&2; exit 1; }

run() {
  local cmd="$*"
  if $DRY_RUN; then
    printf "  \033[90m[dry-run]\033[0m %s\n" "$cmd"
  else
    printf "  \033[90m\$\033[0m %s\n" "$cmd"
    eval "$cmd"
  fi
}

ensure_upstream_remote() {
  local upstream_url="https://github.com/zitadel/zitadel.git"
  local existing_url=""

  if existing_url="$(git remote get-url upstream 2>/dev/null)"; then
    # Normalise for comparison: strip .git, accept both https and ssh forms
    local norm="${existing_url%.git}"
    norm="${norm#https://github.com/}"
    norm="${norm#git@github.com:}"
    if [[ "$norm" != "zitadel/zitadel" ]]; then
      fail "Remote 'upstream' points to '$existing_url' (expected zitadel/zitadel). Fix with: git remote set-url upstream $upstream_url"
    fi
    return 0
  fi

  if $DRY_RUN; then
    warn "Upstream remote not configured. In a real run would: git remote add upstream $upstream_url"
    warn "Dry-run cannot validate the tag without the remote — re-run without --dry-run, or configure upstream manually first."
    UPSTREAM_DRY_DEFERRED=true
  else
    say "Configuring 'upstream' remote -> $upstream_url (first-time setup)"
    git remote add upstream "$upstream_url"
  fi
}

ensure_branch_for_upstream() {
  local upstream_tag="$1"

  if [[ ! "$upstream_tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    fail "Invalid --branch '$upstream_tag'. Expected an upstream tag like vX.Y.Z (e.g. v4.15.0)."
  fi

  ensure_upstream_remote

  if [[ "${UPSTREAM_DRY_DEFERRED:-false}" == "true" ]]; then
    warn "(dry-run) skipping upstream tag validation for '$upstream_tag'"
  else
    say "Verifying upstream tag '$upstream_tag' exists"
    if ! git ls-remote --tags --exit-code upstream "refs/tags/$upstream_tag" >/dev/null 2>&1; then
      fail "Tag '$upstream_tag' not found in upstream. Verify the version name at https://github.com/zitadel/zitadel/tags"
    fi
  fi

  local target_branch="${upstream_tag}-x"
  local current
  current="$(git rev-parse --abbrev-ref HEAD)"

  if [[ "$current" == "$target_branch" ]]; then
    say "Already on fork branch '$target_branch'"
    return 0
  fi

  if ! git diff --quiet || ! git diff --staged --quiet; then
    fail "Cannot switch to '$target_branch': working tree is not clean. Commit or stash first."
  fi

  if git show-ref --verify --quiet "refs/heads/$target_branch"; then
    say "Switching to existing local branch '$target_branch'"
    run "git checkout '$target_branch'"
    return 0
  fi

  if git fetch origin "$target_branch" 2>/dev/null \
     && git show-ref --verify --quiet "refs/remotes/origin/$target_branch"; then
    say "Checking out '$target_branch' from origin"
    run "git checkout -t 'origin/$target_branch'"
    return 0
  fi

  warn ""
  warn "Fork branch '$target_branch' does not exist locally or on origin."
  warn "Will bootstrap it from upstream tag '$upstream_tag'."
  warn ""
  warn "IMPORTANT: The new branch will contain NO Socrate cherry-picks. You must"
  warn "cherry-pick the Socrate-authored commits from the previous patch branch"
  warn "(see patched-zitadel-release.md) BEFORE building a release."

  if ! $YES && ! $DRY_RUN; then
    printf "\n\033[1mCreate new fork branch '%s' from upstream tag '%s'? [y/N]\033[0m " \
      "$target_branch" "$upstream_tag"
    read -r reply
    case "$reply" in y|Y|yes|YES) ;; *) fail "Aborted by user." ;; esac
  fi

  say "Fetching upstream tag '$upstream_tag'"
  run "git fetch upstream 'refs/tags/$upstream_tag:refs/tags/$upstream_tag'"

  say "Creating branch '$target_branch' from tag '$upstream_tag'"
  run "git checkout -b '$target_branch' '$upstream_tag'"

  printf "\n\033[33m! Bootstrap complete — release deferred.\033[0m\n\n"
  cat <<EOF
Next steps:
  1. List the Socrate commits on the previous patch branch:
       git log --author='Alexis KINSELLA' --reverse --oneline <previous-branch>
  2. Cherry-pick them in order; resolve conflicts as you go:
       git cherry-pick <sha>
  3. Push the new branch:
       git push -u origin $target_branch
  4. Re-run this script (drop --branch) to build and release.
EOF
  exit 0
}

preflight() {
  cd "$REPO_ROOT"

  if [[ -n "$UPSTREAM_TAG" ]]; then
    ensure_branch_for_upstream "$UPSTREAM_TAG"
  fi

  local branch
  branch="$(git rev-parse --abbrev-ref HEAD)"
  if [[ ! "$branch" =~ ^v[0-9]+\.[0-9]+\.[0-9]+-[xX]$ ]]; then
    fail "Not on a patch branch. Current: '$branch'. Expected: vX.Y.Z-x (e.g. v4.7.6-x)."
  fi
  CURRENT_BRANCH="$branch"
  say "Branch: $branch"

  if ! git diff --quiet || ! git diff --staged --quiet; then
    fail "Working tree is not clean. Commit or stash changes before releasing."
  fi

  for tool in docker pnpm corepack jq curl base64 sed; do
    command -v "$tool" >/dev/null 2>&1 \
      || fail "Required tool not in PATH: $tool. Are you inside the Dev Container?"
  done

  if ! docker info >/dev/null 2>&1; then
    fail "docker daemon unreachable. Start Docker Desktop and retry."
  fi

  UPSTREAM_BASE="${branch#v}"
  UPSTREAM_BASE="${UPSTREAM_BASE%-[xX]}"
  say "Upstream base: $UPSTREAM_BASE"
}

fetch_scaleway_tags() {
  local creds_b64 user pass token scope
  creds_b64="$(jq -r --arg reg "$SCALEWAY_REGISTRY" '.auths[$reg].auth // empty' "$HOME/.docker/config.json" 2>/dev/null || true)"
  if [[ -z "$creds_b64" ]]; then
    warn "No credentials for $SCALEWAY_REGISTRY in ~/.docker/config.json. Run: docker login $SCALEWAY_REGISTRY"
    return 1
  fi
  local creds
  creds="$(printf '%s' "$creds_b64" | base64 -d)"
  user="${creds%%:*}"
  pass="${creds#*:}"

  scope="repository:${SCALEWAY_PROBE_NAMESPACE}/zitadel:pull"
  token="$(curl -fsSL -u "$user:$pass" \
    "https://${SCALEWAY_REGISTRY}/auth?service=${SCALEWAY_REGISTRY}&scope=${scope}" \
    | jq -r '.token // .access_token // empty')" || return 1

  if [[ -z "$token" ]]; then
    warn "Failed to obtain Scaleway bearer token."
    return 1
  fi

  curl -fsSL -H "Authorization: Bearer $token" \
    "https://${SCALEWAY_REGISTRY}/v2/${SCALEWAY_PROBE_NAMESPACE}/zitadel/tags/list" \
    | jq -r '.tags[]? // empty'
}

fetch_local_tags() {
  curl -fsSL "http://${LOCAL_REGISTRY}/v2/zitadel/tags/list" 2>/dev/null \
    | jq -r '.tags[]? // empty' \
    || true
}

discover_version() {
  if [[ -n "$VERSION_OVERRIDE" ]]; then
    [[ "$VERSION_OVERRIDE" =~ ^v[0-9]+\.[0-9]+\.[0-9]+-[0-9]+$ ]] \
      || fail "Invalid --version '$VERSION_OVERRIDE'. Expected vX.Y.Z-N."
    RESOLVED_VERSION="$VERSION_OVERRIDE"
    say "Using override version: $RESOLVED_VERSION"
    return
  fi

  local tags=""
  case "$TARGET" in
    scaleway) tags="$(fetch_scaleway_tags || true)" ;;
    local)    tags="$(fetch_local_tags    || true)" ;;
  esac

  local escaped_base="${UPSTREAM_BASE//./\\.}"
  local max=0 n
  while IFS= read -r tag; do
    [[ -z "$tag" ]] && continue
    if [[ "$tag" =~ ^v${escaped_base}-([0-9]+)$ ]]; then
      n="${BASH_REMATCH[1]}"
      (( n > max )) && max="$n"
    fi
  done <<<"$tags"

  if (( max == 0 )); then
    warn "No existing tags matching v${UPSTREAM_BASE}-N at the $TARGET registry."
    warn "Defaulting to v${UPSTREAM_BASE}-1 (first patch on this branch)."
    RESOLVED_VERSION="v${UPSTREAM_BASE}-1"
  else
    RESOLVED_VERSION="v${UPSTREAM_BASE}-$((max + 1))"
    say "Highest tag found: v${UPSTREAM_BASE}-${max} - next: $RESOLVED_VERSION"
  fi
}

confirm() {
  printf "\n\033[1m-> Release %s to %s registry. Continue? [y/N]\033[0m " \
    "$RESOLVED_VERSION" "$TARGET"
  if $YES || $DRY_RUN; then
    printf "(auto-confirmed: %s)\n\n" "$([[ $YES == true ]] && echo --yes || echo --dry-run)"
    return 0
  fi
  read -r reply
  case "$reply" in
    y|Y|yes|YES) ;;
    *) fail "Aborted by user." ;;
  esac
}

do_reset() {
  if $SKIP_RESET; then
    say "Skipping workspace reset (--skip-reset)."
    return
  fi
  say "Resetting workspace"
  run "rm -rf node_modules .artifacts/* apps/login/.next apps/login/node_modules"
}

do_install() {
  say "Installing dependencies"
  run "corepack enable"
  run "pnpm install"
  run "pnpm --filter ./apps/login install"
}

do_build() {
  say "Building (pnpm nx pack)"
  run "pnpm nx clean"
  run "pnpm nx pack"
}

do_push() {
  case "$TARGET" in
    scaleway)
      say "Tagging and pushing $RESOLVED_VERSION to ${SCALEWAY_NAMESPACES[*]}"
      run "VERSION=$RESOLVED_VERSION bash scripts/docker-tag-push.sh"
      ;;
    local)
      say "Tagging and pushing $RESOLVED_VERSION to $LOCAL_REGISTRY"
      run "VERSION=$RESOLVED_VERSION bash scripts/docker-tag-push-local.sh"
      ;;
  esac
}

summary() {
  printf "\n\033[32mok Release complete: %s -> %s\033[0m\n\n" "$RESOLVED_VERSION" "$TARGET"
  case "$TARGET" in
    local)
      cat <<EOF
Next steps (smoke test):
  1. In socrate-k8s, TEMPORARILY set image.tag: $RESOLVED_VERSION in
     manifests/zitadel/values.yaml
  2. just deploy-zitadel
  3. Open https://zitadel.internal.local.socrate.ninja and verify login
  4. REVERT manifests/zitadel/values.yaml -- socrate-local must stay on
     the unpatched upstream tag.
EOF
      ;;
    scaleway)
      cat <<EOF
Next steps (environment rollout via socrate-flux):
  Bump image.tag to $RESOLVED_VERSION in the target environment's values
  file, commit, and let Flux reconcile. Promote progressively:

    dev / qa: socrate/sandbox/qa/socrate/zitadel/values.yaml
    staging : socrate/sandbox/staging/socrate/zitadel/values.yaml
    prod    : socrate/prod/prod/socrate/zitadel/values.yaml
    OVH PRA : socrate/ovh/prod/socrate/zitadel/values.yaml
              (verify image.repository -- may need a separate mirror push)

  Roll to staging first, observe one login cycle, then promote to prod.
EOF
      ;;
  esac
  printf "\n"
}

preflight
discover_version
confirm
do_reset
do_install
do_build
do_push
summary
