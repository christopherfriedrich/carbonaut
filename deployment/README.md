# Carbonaut Infrastructure Deployment

## Carbonaut Infrastructure Architecture

...

## Using the Helm chart

...

## Using the gRPC Envoy Proxy locally

To enable communication between server and frontend using gRPC, an [envoy proxy](https://www.envoyproxy.io) must be running and connected to the frontend. 

 To run envoy, you must install the tool locally (see the [installation instructions](https://www.envoyproxy.io/docs/envoy/latest/start/start)) and start the proxy by specifying the configuration file - command: `envoy -c deployment/grpc-web-envoy.yaml`.

* Localhost admin console: `http://localhost:9901` 
* Endpoint to route gRPC: `http://localhost:8080`

