# Carbonaut API

Carbonaut uses gRPC for communication and therefore uses a `.proto` file to define the API between the **Carbonaut-UI** and the **Carbonaut-API**. The file `carbonaut.proto` defines the API and is being used to generate the configuration for typescript (used by the carbonaut-ui) and golang (used by the carbonaut-api).

## Generate gRPC Server and Client configuration files

Required installs:

1. `protoc`: [Installation](https://grpc.io/docs/protoc-installation/) 
   1. **NOTE** as of 18th of October 2022 the latest `protobuf` versions has a bug which does not allow us to generate typescript code. Use the previous previous version. Macos: `$brew install protobuf@3 && brew link --overwrite protobuf@3`.
2. `grpc-web`: [Node Package](https://www.npmjs.com/package/grpc-web) example command `npm i -g grpc-web`
3. `protoc-gen-go`: [Go Package](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go) example command `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
4. `protoc-gen-go-grpc`: [Go Package](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc) example command `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

To generate client and server configurations use the Makefile command `make compile-grpc`. You should see file being generated under `ui/packages/api/dist/*` and `pkg/api/v*`.
