#!/bin/bash
# scripts/docker-tag-push.sh

VERSION=${VERSION:-"latest"}

# Copy the Nx-built binary to the expected name
cp .artifacts/bin/$(go env GOOS)/$(go env GOARCH)/zitadel.local zitadel

# Build Docker image with multiple tags
DOCKER_BUILDKIT=1 docker build \
  --no-cache \
  --platform linux/amd64 \
  --build-arg TARGETPLATFORM=linux/amd64 \
  --build-arg BUILDPLATFORM=linux/amd64 \
  --build-arg GOOS=linux \
  --build-arg GOARCH=amd64 \
  -f apps/api/Dockerfile \
  -t rg.fr-par.scw.cloud/cr-registry-dev/zitadel:${VERSION} \
  -t rg.fr-par.scw.cloud/cr-registry-qa/zitadel:${VERSION} \
  -t rg.fr-par.scw.cloud/cr-registry-staging/zitadel:${VERSION} \
  -t rg.fr-par.scw.cloud/cr-registry-prod/zitadel:${VERSION} .

# Push to all registries
docker push rg.fr-par.scw.cloud/cr-registry-dev/zitadel:${VERSION}
docker push rg.fr-par.scw.cloud/cr-registry-qa/zitadel:${VERSION}
docker push rg.fr-par.scw.cloud/cr-registry-staging/zitadel:${VERSION}
docker push rg.fr-par.scw.cloud/cr-registry-prod/zitadel:${VERSION}