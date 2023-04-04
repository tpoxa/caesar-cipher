Technical challenge 
========================
Caesar cipher gRPC server and client

Prerequisites
========================
Install golang and protoc compiler with gRPC plugin (https://grpc.io/docs/languages/go/quickstart/#prerequisites)

Build
========================
```shell
make build
```

Run server
========================
Default listen addr is localhost:8080
```shell
cd bin/
./server
./server --help #for options
```

Client usage
========================
```shell
./client encrypt maksym
./client decrypt nbltzn
./client --help # for options
```
