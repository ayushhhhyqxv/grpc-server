
# GRPC-Server

Project demonstrates different types of gRPC communication patterns implemented in Go:

- **Unary RPC**
- **Server Streaming**
- **Client Streaming**
- **Bidirectional Streaming**

Repository is structured into **client**, **server**, and **proto** directories.


---

## Setup

### Install Dependencies
Make sure you have **Go**, **Protocol Buffers**, and **gRPC for Go** installed.


```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
````



Set PATH !
Add `$GOPATH/bin` to your `PATH` if not already.




### Run Server

Navigate to the `server` folder and start the gRPC server:

```bash
cd server
go run *.go
```



### Run Client

In a separate terminal, start the client:

```bash
cd client
go run *.go
```



## Contains:

* Unary RPC calls
* Server-side streaming
* Client-side streaming
* Bidirectional streaming


---


## Requirements

* Go 1.18+
* gRPC-Go
* protocol Buffers compiler (`protoc`)

---




