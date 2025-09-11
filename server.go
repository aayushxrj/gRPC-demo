package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aayushxrj/gRPC-demo/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer 
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error){
	return &pb.AddResponse{
		Sum : req.A + req.B,
	}, nil
}

func main (){

	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(grpcServer, &server{})


	log.Printf("Server is running on the port%s", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}