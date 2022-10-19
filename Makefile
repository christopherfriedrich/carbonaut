# Copyright (c) 2022 CARBONAUT AUTHOR
#
# Licensed under the MIT license: https://opensource.org/licenses/MIT
# Permission is granted to use, copy, modify, and redistribute the work.
# Full license information available in the project LICENSE file.

verify: verify-go-mod verify-git verify-build verify-lint verify-test-unit

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

install-go: 
	go get ./...

compile-grpc:
	protoc -I=api/v1 --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative api/v1/api.proto
	protoc -I=api/v1 --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative api/v1/*/*.proto

	mkdir -p ui/packages/@carbonaut-cloud-api/dist

	protoc -I=api/v1 --js_out=import_style=commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist --grpc-web_out=import_style=typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist api/v1/api.proto
	protoc -I=api/v1 --js_out=import_style=commonjs,binary:./ui/packages/@carbonaut-cloud-api/dist --grpc-web_out=import_style=typescript,mode=grpcweb:./ui/packages/@carbonaut-cloud-api/dist api/v1/*/*.proto

    # Workaround to make generated files usable as dependency
	echo "export * from './ApiServiceClientPb';\nexport * from './api_pb';\r" > ui/packages/@carbonaut-cloud-api/dist/index.ts
