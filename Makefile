# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

verify: check-misc check-go check-ts check-go-e2e

check-misc: check-git
check-go: check-go-mod check-go-build check-go-lint
check-go-e2e: check-go-test-unit
check-ts: check-ts-lint

# GO check's
check-go-build:
	./hack/check-go-build.sh

check-go-lint:
	./hack/check-go-lint.sh

check-go-test-unit:
	./hack/check-go-test.sh

check-go-mod:
	go vet ./...
	go mod tidy -compat=1.18

# TYPESCRIPT check's
check-ts-lint:
	echo "todo"

# MISC check's
check-git:
	git diff --exit-code

#
# Commands to install/setup the environment and generate artifacts
#

# NOTE: it's required to install grpc libraries first, see file: carbonaut/api/v1/README.md
install: compile-grpc install-go install-pnpm

install-pnpm:
	echo "install-pnpm steps not implemented; TODO"

install-go:
	go get ./...

compile-grpc:
	protoc -I=api/v1 --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative api/v1/api.proto
	protoc -I=api/v1 --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative api/v1/*/*.proto

	mkdir -p ui/packages/@carbonaut-cloud-api/dist

	protoc -I=api/v1 --js_out=import_style=commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist --grpc-web_out=import_style=typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist api/v1/api.proto
	protoc -I=api/v1 --js_out=import_style=commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist --grpc-web_out=import_style=typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist api/v1/*/*.proto

    # NOTE: Workaround to make generated files usable as dependency
	echo "export * from './ApiServiceClientPb';\nexport * from './api_pb';\r" > ui/packages/@carbonaut-cloud-api/dist/index.ts

#
# Commands for development
#

upgrade-go:
	go get -u -t ./...

run-api-server-with-fake:
	go run cmd/api/main.go -fake-data -port 50051
