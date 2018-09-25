package main

import (
	"github.com/iron-kit/go-ironic"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/services/core/handler"
	"iunite.club/services/core/subscriber"

	example "iunite.club/services/core/proto/example"
)

func main() {
	// New Service
	service := ironic.NewService(
		micro.Name("iunite.club.srv.core"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
