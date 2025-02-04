FROM ubuntu:22.04

# Set build arguments with default values
ARG NODE_VERSION=20
ARG GO_VERSION=1.20

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl gnupg2 lsb-release software-properties-common

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION}.x | bash -
RUN apt-get install -y nodejs

# Install Go
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzf - -C /usr/local
ENV PATH="/usr/local/go/bin:$PATH"

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./
RUN pnpm install

ENTRYPOINT tail -f /dev/null