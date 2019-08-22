package main

import (
	"context"
	"fmt"
	pb "github.com/gologic/microservice/consignment-service/consignment"
	"github.com/gologic/microservice/vessel-service/proto/vessel"
	userService "github.com/gologic/microservice/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
	"os"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/client"
	"errors"
)

const (
	defaultHost = "datastore:27017"
)

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token:", token)

		// Auth here
		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
func main() {
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
		micro.WrapHandler(AuthWrapper),
	)
	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())
	consignmentCollection := client.Database("shippy").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}

	vesselClient := vessel.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
