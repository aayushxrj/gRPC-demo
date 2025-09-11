
```
go mod init github.com/aayushxrj/gRPC-demo
go get google.golang.org/grpc
```

```
protoc --go_out=. --go-grpc_out=. proto/main.proto
```

