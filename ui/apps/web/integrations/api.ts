import * as grpcWeb from 'grpc-web';
import { GreeterClient, HelloRequest, HelloReply } from '@carbonaut-cloud/api';

const service = new GreeterClient('http://localhost:8080', null, null);

const request = new HelloRequest();
request.setName('John');

const call = service.sayHello(
  request,
  { 'custom-header-1': 'value1' },
  (error: grpcWeb.RpcError, response: HelloReply) => {
    console.log(response.getMessage());
  }
);

call.on('status', (status: grpcWeb.Status) => {
  // ...
});
