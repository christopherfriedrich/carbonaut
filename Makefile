# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

verify: swag verify-go-mod verify-git verify-build verify-lint verify-test-unit

verify-build:
	./hack/check-build.sh

verify-lint:
	./hack/check-lint.sh

verify-test-unit:
	./hack/check-test.sh

verify-go-mod:
	go vet ./...
	go mod tidy -compat=1.18

verify-git:
	git diff --exit-code

upgrade:
	go get -u -t ./...

install: 
	# install swagger tool to compile swagger carbonaut api definition 
	go install github.com/swaggo/swag/cmd/swag@v1.8.6
	go get ./...

compile-grpc:
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/v1/carbonaut.proto

	mkdir -p ui/packages/api/dist

	protoc api/v1/carbonaut.proto --js_out=import_style=commonjs,binary:./ui/packages/api/dist --grpc-web_out=import_style=typescript,mode=grpcweb:./ui/packages/api/dist

	mv ui/packages/api/dist/api/v1/* ui/packages/api/dist
	rmdir ui/packages/api/dist/api/v1 ui/packages/api/dist/api
