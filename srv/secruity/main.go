package main

import (
	"github.com/iron-kit/unite-services/secruity/client"
	"github.com/iron-kit/unite-services/secruity/handler"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	auth "github.com/iron-kit/unite-services/secruity/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.srv.secruity"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.WrapHandler(
			client.UserWrapper(service),
		),
	)

	// Register Handler
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.secruity", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.secruity", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
