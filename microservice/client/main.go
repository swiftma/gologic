package main

import (
	"io/ioutil"

	"context"

	pb "github.com/gologic/microservice/consignment-service/consignment"

	"encoding/json"
	"github.com/micro/go-micro"
	"log"
	"errors"
	"os"
	"github.com/micro/go-micro/metadata"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.service.consignment"))
	service.Init()

	client := pb.NewShippingServiceClient("shippy.service.consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}
	file = os.Args[1]
	token := os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	str, _ := json.MarshalIndent(getAll.Consignments,"","  ")
	log.Printf("%s", str)

}
