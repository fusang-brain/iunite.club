package main

import (
	"github.com/iron-kit/unite-services/api-gateway/handler"
	auth "github.com/iron-kit/unite-services/api-gateway/proto/auth"
	"github.com/micro/go-api"
	"github.com/micro/go-api/handler/rpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.api.api-gateway"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "Auth.Login",
				Path:    []string{"/auth/login"},
				Method:  []string{"POST"},
				Handler: rpc.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "Auth.Register",
				Path:    []string{"/auth/register"},
				Method:  []string{"POST"},
				Handler: rpc.Handler,
			},
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
