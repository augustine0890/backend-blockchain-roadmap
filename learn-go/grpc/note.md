# Go Protocol Buffer
## Introduction
- __REST__: the resource-oriented architecture (ROA)-style distributed application.
- __GraphQL__: a single endpoint that you can use to query or mutate the data through HTTP verbs.
- __WebSockets__: support bi-direactional communication over a single TCP connection.
- __Server Sent Eventing__: the server sends messages to the client once the initial connection has been set up by the client.
- __gRPC__
  - The idea behind gRPC is to enable developers to use RPC-like communications over HTTP/2 while have a single client library.
  
## Protocol Buffer Data Format
- Protocol Buffers are a data format, much like JSON or XML in the sense that they store structured data which can be serialized or deserialized.
- Smaller than XML or JSON.
- Simple message definition:
  ```proto
  message Test1 {
    optional int32 a = 1;
  }
  ```
- Encoded message:
  `08 96 01`

## Packages
- Mac OS 
  - `brew install protobuf`
- `go get google.golang.org/protobuf`
- Install the Go protocol buffers plugin:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- Compile
  - `protoc --go-grpc_out=. *.proto`


## References
- Protocol Buffers by [Google](https://developers.google.com/protocol-buffers/docs/overview) Developers
- Introduction to gRPC [blog](https://sahansera.dev/introduction-to-grpc/)