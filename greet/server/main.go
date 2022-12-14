package main

import (
	pb "grpc-go-course/greet/proto"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
