package main

import (
	"github.com/iron-kit/go-ironic"
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"iunite.club/services/navo/client"
	"iunite.club/services/navo/handler"
	auth "iunite.club/services/navo/proto/auth"
	// example "iunite.club/services/navo/proto/example"
)

func main() {
	// New Service
	service := ironic.NewService(
		micro.Name("iunite.club.api.navo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		// micro.WrapHandler(client.ExampleWrapper(service)),
		micro.WrapHandler(client.SecruityWrapper(service)),
	)

	// Register Handler
	auth.RegisterAuthHandlerHandler(service.Server(), new(handler.AuthHandler))
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
