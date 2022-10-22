#!/usr/bin/env bash

# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

# Exit script if you try to use an uninitialized variable.
set -o nounset
# Exit script if a statement returns a non-true return value.
set -o errexit
# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

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

