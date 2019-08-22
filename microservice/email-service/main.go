package main

import (
	"log"

	pb "github.com/gologic/microservice/user-service/proto/user"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"context"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
		micro.Version("latest"),
	)

	srv.Init()
	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	//// Get the broker instance using our environment variables
	//pubsub := srv.Server().Options().Broker
	//if err := pubsub.Connect(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Subscribe to messages on the broker
	//_, err := pubsub.Subscribe(topic, func(p broker.Event) error {
	//	var user *pb.User
	//	if err := json.Unmarshal(p.Message().Body, &user); err != nil {
	//		return err
	//	}
	//	log.Println(user)
	//	go sendEmail(user)
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Println(err)
	//}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
//
//func sendEmail(user *pb.User) error {
//	log.Println("Sending email to:", user.Name)
//	return nil
//}
