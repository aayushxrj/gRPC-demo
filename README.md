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
# TLS

old .X509 certificate doesn't work anymore

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config cert.conf  
```