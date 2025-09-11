package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aayushxrj/gRPC-demo/proto/gen"
	"github.com/aayushxrj/gRPC-demo/proto/gen/farewell"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer 
	farewellpb.UnimplementedGoodByeServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error){
	log.Println("Add RPC called.")
	return &pb.AddResponse{
		Sum : req.A + req.B,
	}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error){
	log.Println("Greet RPC called.")
	return &pb.HelloResponse{
		Message: "Hello "+req.GetName(),
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error){
	log.Println("Farewell RPC called.")
	return &farewellpb.GoodByeResponse{
		Message: "GoodBye "+req.GetName(),
	}, nil
}

func main (){

	port := ":50051"
	cert := "cert.pem"
	key := "key.pem"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Failed to load credentials:", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})
	farewellpb.RegisterGoodByeServer(grpcServer, &server{})

	log.Printf("Server is running on the port%s", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}