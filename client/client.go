package main

import (
	"context"
	"log"
	"time"

	mainapipb "github.com/aayushxrj/gRPC-demo/proto/gen"
	farewellpb "github.com/aayushxrj/gRPC-demo/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/credentials/insecure"
)

func main() {

	port := ":50051"
	cert := "cert.pem"

	// for without tls connections
	// conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Println("Unable to connet", err)
	// }
	// defer conn.Close()

	// with TLS
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatal("Failed to load credentials:", err)
	}

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println("Unable to connet", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)
	client2 := mainapipb.NewGreeterClient(conn)
	client3 := farewellpb.NewGoodByeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &mainapipb.AddRequest{
		A : 10,
		B: 20,
	}

	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	req2 := &mainapipb.HelloRequest{
		Name : "Aayush",
	}
	res2, err := client2.Greet(ctx, req2)
	if err != nil {
		log.Fatalln("Could not greet", err)
	}

	req3 := &farewellpb.GoodByeRequest{
		Name : "Aayush",
	}
	res3, err := client3.BidGoodBye(ctx, req3)
	if err != nil {
		log.Fatalln("Could not greet", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("Greet:", res2.Message)
	log.Println("Farewell:", res3.Message)

	state := conn.GetState()
	log.Println("Connection state:",state)
}
