# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

verify: compile-grpc check-misc check-go check-ts check-go-e2e

check-misc: check-git
check-go: check-go-mod check-go-build check-go-lint
check-go-e2e: check-go-test-unit
check-ts: check-ts-lint

# GO check's
check-go-build:
	./hack/check-go-build.bash

check-go-lint:
	./hack/check-go-lint.bash

check-go-test-unit:
	./hack/check-go-test.bash

check-go-mod:
	go vet ./...
	go mod tidy -compat=1.19

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
	./hack/compile-grpc.bash

#
# Commands for development
#

upgrade-go:
	go get -u -t ./...

run-api-server-with-fake:
	go run cmd/api/main.go -fake-data -port 50051

#
# DEPLOY CONTAINER IMAGE
build-and-push:
	./hack/push-and-deploy.bash agent
	./hack/push-and-deploy.bash api
