package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/iron-kit/unite-services/oa-srv/handler"
	"github.com/iron-kit/unite-services/oa-srv/subscriber"

	example "github.com/iron-kit/unite-services/oa-srv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.srv.srv.oa-srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("kit.iron.srv.srv.oa-srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("kit.iron.srv.srv.oa-srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
