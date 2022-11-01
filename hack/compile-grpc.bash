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

_log() {
    echo " $(date +'[%F %T]') - $1"
}

# Required: authentication to container registry

_log "Generate gRPC files for carbonaut-api"

protoc -I=api\/v1 --go_out=pkg\/api --go_opt=paths=source_relative --go-grpc_out=pkg\/api --go-grpc_opt=paths=source_relative api\/v1\/api.proto

protoc -I=api\/v1 --go_out=pkg\/api --go_opt=paths=source_relative --go-grpc_out=pkg\/api --go-grpc_opt=paths=source_relative api\/v1\/*\/*.proto

mkdir -p ui\/packages\/@carbonaut-cloud-api\/dist

_log "Generate gRPC files for carbonaut-ui"

protoc -I="api/v1" --js_out="import_style"="commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist" --grpc-web_out="import_style=typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist" api/v1/api.proto
protoc -I="api/v1" --js_out="import_style=commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist" --grpc-web_out="import_style"="typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist" api/v1/*/*.proto

# NOTE: Workaround to make generated files usable as dependency
printf "export * from './ApiServiceClientPb';\nexport * from './api_pb';\r" > ui/packages/@carbonaut-cloud-api/dist/index.ts
