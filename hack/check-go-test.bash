#!/bin/bash

# Copyright (c) 2022 CARBONAUT AUTHORS
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

set -euo pipefail

TEST_TIMEOUT=800

for arg in "$@"
do
    case $arg in
        -t=*|--timeout=*)
        TEST_TIMEOUT="${arg#*=}"
        shift
        ;;
        -t|--timeout)
        TEST_TIMEOUT="$2"
        shift
        shift
    esac
done

# REPO_ROOT=$(git rev-parse --show-toplevel)

GO111MODULE=on go test -timeout="${TEST_TIMEOUT}s" -count=1 -cover -coverprofile coverage.out $(go list ./... )
go tool cover -html coverage.out -o coverage.html

