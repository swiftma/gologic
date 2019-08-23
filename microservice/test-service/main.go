package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/gologic/microservice/test-service/handler"
	"github.com/gologic/microservice/test-service/subscriber"

	test "github.com/gologic/microservice/test-service/proto/test"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("hello.test"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test", service.Server(), new(subscriber.Test))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
