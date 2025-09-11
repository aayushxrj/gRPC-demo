# Server
```
go mod init github.com/aayushxrj/gRPC-demo
go get google.golang.org/grpc
go mod tidy
```
```
protoc --go_out=. --go-grpc_out=. proto/main.proto
```

# Client 
Generate code again as the client only has the main.proto file 

```
protoc --go_out=. --go-grpc_out=. proto/main.proto
```
```
go mod init github.com/aayushxrj/gRPC-demo
go get google.golang.org/grpc
go mod tidy
```
