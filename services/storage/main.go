package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/services/storage/handler"
	"iunite.club/services/storage/subscriber"

	example "iunite.club/services/storage/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("iunite.club.srv.storage"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.storage", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.storage", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
