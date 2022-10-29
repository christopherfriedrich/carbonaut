#!/usr/bin/env bash

# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

set -o errexit
set -o pipefail
set -o nounset

PLATFORMS=(
    linux/amd64
    linux/386
    linux/arm
    linux/arm64
    linux/ppc64le
    linux/s390x
    windows/amd64
    windows/386
    freebsd/amd64
    darwin/amd64
)

for PLATFORM in "${PLATFORMS[@]}"; do
    OS="${PLATFORM%/*}"
    ARCH=$(basename "$PLATFORM")

    echo "Build on platform $PLATFORM"
    GOARCH="$ARCH" GOOS="$OS" go build ./...
done

