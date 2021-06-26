package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "tungnguyen.shippy/shippy-service-consignment/proto/consignment"
	"tungnguyen.shippy/shippy-service-consignment/repository"
	srv "tungnguyen.shippy/shippy-service-consignment/service"
)

const (
	port = ":50051"
)

func main() {
	repo := repository.NewRepo()
	srv := srv.NewService(srv.ServiceConfig{Repo: repo})

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cann't establish network : %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, srv)
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
