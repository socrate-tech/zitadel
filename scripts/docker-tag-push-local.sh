#!/bin/bash
# scripts/docker-tag-push.sh

VERSION=${VERSION:-"latest"}

# Copy the Nx-built binary to the expected name
cp .artifacts/bin/$(go env GOOS)/$(go env GOARCH)/zitadel.local zitadel

# Build Docker image with multiple tags
DOCKER_BUILDKIT=1 docker build \
  --no-cache \
  -f apps/api/Dockerfile \
  -t registry.internal.local.socrate.ninja:5001/zitadel:${VERSION} .

# Push to all registries
docker push registry.internal.local.socrate.ninja:5001/zitadel:${VERSION}