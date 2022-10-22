#!/usr/bin/env bash

# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

set -euo pipefail

# Required: authentication to container registry

# Container registry used to 
REGISTRY=ghcr.io
ORG=carbonaut-cloud

# Name of the container to build and deploy, e.g.: agent, analytics, api, ui
CONTAINER=$1

echo ":: Build container image $ORG/$CONTAINER ::"
SHORTHASH="$(git rev-parse --short HEAD)"
docker build -f build/Containerfile.$CONTAINER -t $REGISTRY/$ORG/carbonaut-$CONTAINER:latest -t $REGISTRY/$ORG/carbonaut-$CONTAINER:$SHORTHASH .
echo ":: Push container image $REGISTRY/$ORG/carbonaut-$CONTAINER:latest:$SHORTHASH to $REGISTRY ::"
docker push $REGISTRY/$ORG/carbonaut-$CONTAINER:latest
docker push $REGISTRY/$ORG/carbonaut-$CONTAINER:$SHORTHASH
