package main

import (
	"log"
	"os"

	pb "github.com/gologic/microservice/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/micro/cli"
	//"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro"
)


func main() {

	//cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("shippy.auth", microclient.DefaultClient)

	// Define our flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name: "company",
				Usage: "Your company",
			},
		),
	)

	// Start as service
	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			authResponse, err := client.Auth(context.TODO(), &pb.User{
				Email:    email,
				Password: password,
			})

			if err != nil {
				log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
			}

			log.Printf("Your access token is: %s \n", authResponse.Token)

			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}