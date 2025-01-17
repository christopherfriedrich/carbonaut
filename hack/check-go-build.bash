#!/bin/bash

# Copyright (c) 2022 CARBONAUT AUTHORS
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

# This is the concurrency limit
MAX_POOL_SIZE=3
# This is used within the program. Do not change.
CURRENT_POOL_SIZE=0

# Jobs will be loaded from this file
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

    # This is the blocking loop where it makes the program to wait if the job pool is full
    while [ $CURRENT_POOL_SIZE -ge $MAX_POOL_SIZE ]; do
        CURRENT_POOL_SIZE=$(jobs | wc -l)
    done
    echo " $(date +'[%F %T]') - Build on platform $PLATFORM"
    GOARCH="$ARCH" GOOS="$OS" go build ./... &
    CURRENT_POOL_SIZE=$(jobs | wc -l)
done

# wait for all background jobs (forks) to exit before exiting the parent process
wait
