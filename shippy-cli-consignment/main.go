package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	pb "tungnguyen.shippy/shippy-service-consignment/proto/consignment"
)

const (
	addr            = "localhost:50051"
	defaultFilePath = "consignment.json"
)

func parseFile(filePath string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	consignment := pb.Consignment{}
	if err = json.Unmarshal(data, &consignment); err != nil {
		return nil, err
	}

	return &consignment, nil
}

func main() {
	filePath := defaultFilePath
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}
	data, err := parseFile(filePath)
	if err != nil {
		log.Fatalf("Cann't parse file : %v\n", err)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cann't create connect to service : %v\n", err)
	}

	client := pb.NewShippingServiceClient(conn)

	res, err := client.CreateConsignment(context.Background(), data)
	if err != nil {
		log.Fatalf("cann't get service response : %v\n", err)
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Create : %v\nConsignment : %v\n", res.Create, res.Consignment)

	res, err = client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("cann't get service response : %v\n", err)
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Create : %v\nConsignments : %v\n", res.Create, res.Consignments)
}
