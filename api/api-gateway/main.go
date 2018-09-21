package main

import (
	"github.com/micro/go-api"
	apiHandler "github.com/micro/go-api/handler/api"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/api/api-gateway/client"
	"iunite.club/api/api-gateway/handler"
	auth "iunite.club/api/api-gateway/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.api.api-gateway"),
		micro.Version("latest"),
		// micro.WrapClient()
	)

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "Auth.Login",
				Path:    []string{"/auth/login"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "Auth.Register",
				Path:    []string{"/auth/register"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
	)

	// Initialise service
	service.Init(
		micro.WrapHandler(
			client.SecruityWrapper(service),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
