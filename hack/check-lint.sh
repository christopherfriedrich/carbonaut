#!/usr/bin/env bash

# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

set -o errexit
set -o pipefail
set -o nounset

VERSION=v1.50.0
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yaml ]]; then
    echo 'ERROR: missing .golangci.yaml in repo root' >&2
    exit 1
fi

if ! command -v golangci-lint; then
    curl -sfL $URL | sh -s $VERSION
    PATH=$PATH:bin
fi

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
