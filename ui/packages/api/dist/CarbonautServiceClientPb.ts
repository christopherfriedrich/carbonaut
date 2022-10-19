/**
 * @fileoverview gRPC-Web generated client stub for carbonaut.api.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as carbonaut_pb from './carbonaut_pb';


export class GreeterClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSayHello = new grpcWeb.MethodDescriptor(
    '/carbonaut.api.v1.Greeter/SayHello',
    grpcWeb.MethodType.UNARY,
    carbonaut_pb.HelloRequest,
    carbonaut_pb.HelloReply,
    (request: carbonaut_pb.HelloRequest) => {
      return request.serializeBinary();
    },
    carbonaut_pb.HelloReply.deserializeBinary
  );

  sayHello(
    request: carbonaut_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null): Promise<carbonaut_pb.HelloReply>;

  sayHello(
    request: carbonaut_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: carbonaut_pb.HelloReply) => void): grpcWeb.ClientReadableStream<carbonaut_pb.HelloReply>;

  sayHello(
    request: carbonaut_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: carbonaut_pb.HelloReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/carbonaut.api.v1.Greeter/SayHello',
        request,
        metadata || {},
        this.methodInfoSayHello,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/carbonaut.api.v1.Greeter/SayHello',
    request,
    metadata || {},
    this.methodInfoSayHello);
  }

}

