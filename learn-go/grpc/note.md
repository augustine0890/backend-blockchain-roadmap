# Go Protocol Buffer
## Introduction
- __REST__: the resource-oriented architecture (ROA)-style distributed application.
- __GraphQL__: a single endpoint that you can use to query or mutate the data through HTTP verbs.
- __WebSockets__: support bi-direactional communication over a single TCP connection.
- __Server Sent Eventing__: the server sends messages to the client once the initial connection has been set up by the client.
- __gRPC__
  - The idea behind gRPC is to enable developers to use RPC-like communications over HTTP/2 while have a single client library.
- Efficient for inter-process communication
- Code-generation of client and server stubs with strong types.
- Cons:
  - Most web browsers support is limited.
  - Changing the service definition might require rework and regeneration of code.
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

## gRPC Server and Client
- Install protocol buffer compiler, `protoc`
  ```
  $ brew install protobuf
  $ protoc --version  # Ensure compiler version is 3+
  ```
1. Create the definition of the service with a `.proto` file
2. Generate the server-side code in your preferred language (Go, Java, C# etc...)
3. Generate the client-side code in your preferred language --> doesn't have to be the same language with server-side code.
  - This includes methods that we can invoke on our server with additional code to serialize, deserialize messages.
4. The client-server communication happens over an HTTP/2 connection.

## gRPC Modes
- There are 4 modes of gRPC communication styles.
  - Unary RPC: send a request and receive a single response.
  - Server Streaming RPC: client sends a request and reads until the server stops sending messages via a stream
  - Client Streaming: client sends messages through the stream and waits for the server to read and return a response.
  - Bidirectional streaming RPC: both client and server streams messages both ways.

## Packages
- Mac OS 
  - `brew install protobuf`
- `go get google.golang.org/protobuf`
- The Go gRPC plugin
  - `go get google.golang.org/grpc`
- Install the Go protocol buffers plugin:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- Compile
  - `protoc --go-grpc_out=. *.proto`
- Install gRPCurl
  - gRPC Curl can be installed from the following URL:
    - `https://github.com/fullstorydev/grpcurl`
  - Or `go install` command:
    - `go install github.com/fullstorydev/grpcurl/cmd/grpcurl`
# gRPC Server in Go
  ```proto
  syntax = "proto3";
  
  option go_package = "pb/inventory";
  
  message Book {
    string title = 1;
    string author = 2;
    int32 page_count = 3;
    optional string language = 4;
  }
  
  message GetBookListRequest {}
  message GetBookListResponse { repeated Book books = 1;}
  
  service Inventory { 
    rpc GetBookList(GetBookListRequest) returns (GetBookListResponse) {}
  }
  ```
- Generate Go stubs
  - `make gen`
- Clean stubs
  - `make clean`
- The `option` keyword: the Protobuf compiler where we want to put the generted stubs.
- Generate the stubs:
  - `protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.`
- Use `gRPCurl` for invoking RPCs
  - `grpcurl -plaintext localhost:8080 Inventory/GetBookList` # to get a list of books
  - By default gRPCurl will attempt to use server reflection to determine the methods and types. To enable this reflection must be explictly implemented in the API. If reflection is not availabe, gRPCurl can also use the protos for the service.
    - `grpcurl -plaintext localhost:8080 list` # introspect the service
    - `grpcurl -plaintext localhost:8080 list Inventory` # listing methods for service
    - `grpcurl -plaintext localhost:8080 describe Inventory.GetBookList` # methods in details
## References
- Protocol Buffers by [Google](https://developers.google.com/protocol-buffers/docs/overview) Developers
- Introduction to gRPC [blog](https://sahansera.dev/introduction-to-grpc/)