variable "GITHUB_SHA" {
  default = "latest"
}

variable "REGISTRY" {
  default = "ghcr.io/zitadel"
}

group "ci" {
  targets = ["build", "lint", "unit"]
}

group "generate" {
  targets = ["console-generate" , "core-generate"]
}

group "build" {
  targets = ["console-build", "core-build"]
}

group "lint" {
  targets = ["console-lint", "core-lint"]
}

group "unit" {
  targets = ["core-unit"]
}

group "image" {
  targets = ["console-image", "core-image"]
}

target "devcontainer" {
  dockerfile = "Dockerfile.devcontainer"
  context = "."
  tags = ["${REGISTRY}/base:${GITHUB_SHA}"]
  push = false
}

target "_console" {
  dockerfile = "Dockerfile.console"
  context = "."
  contexts = {
    devcontainer = "target:devcontainer"
    nginx = "docker-image://nginx:1.27-alpine"
  }
}

target "console-generate" {
  inherits = ["_console"]
  output = ["type=local,dest=./"]
  target = "generate"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=console-generate"]
  cache-from = ["type=gha,scope=console-generate"]
}

target "console-build" {
  inherits = ["_console"]
  output = ["type=local,dest=.build/console"]
  target = "build"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=console-build"]
  cache-from = ["type=gha,scope=console-build"]
}

target "console-lint" {
  inherits = ["_console"]
  output = ["type=cacheonly"]
  target = "lint"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=console-lint"]
  cache-from = ["type=gha,scope=console-lint"]
}

target "_core" {
  dockerfile = "Dockerfile.core"
  context = "."
  contexts = {
    devcontainer = "target:devcontainer"
  }
}

target "core-generate" {
  inherits = ["_core"]
  output = ["type=local,dest=./"]
  target = "generate"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=core-generate"]
  cache-from = ["type=gha,scope=core-generate"]
}

target "core-build" {
  inherits = ["_core"]
  name = "core-build-${os}-${arch}"
    matrix = {
    os = ["linux", "darwin", "windows"]
    arch = ["amd64", "arm64"]
  }
  args = {
    OS = os
    ARCH = arch
  }
  output = ["type=local,dest=.build/core"]
  contexts = {
    console = "target:console-build"
  }
  target = "build"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=core-build-${os}-${arch}"]
  cache-from = ["type=gha,scope=core-build-${os}-${arch}"]
}

target "core-lint" {
  inherits = ["_core"]
  output = ["type=cacheonly"]
  target = "lint"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=core-lint"]
  cache-from = ["type=gha,scope=core-lint"]
}

target "core-unit" {
  inherits = ["_core"]
  output = ["type=local,dest=.build/core"]
  contexts = {
    console = "target:console-build"
  }
  target = "unit"
  cache-to = ["type=gha,ignore-error=true,mode=max,scope=core-unit"]
  cache-from = ["type=gha,scope=core-unit"]
}