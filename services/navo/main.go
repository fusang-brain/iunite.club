package main

import (
	"github.com/iron-kit/go-ironic"
	"github.com/micro/go-api"
	apiHandler "github.com/micro/go-api/handler/api"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/services/navo/client"
	"iunite.club/services/navo/handler"
	core "iunite.club/services/navo/proto"
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
		micro.WrapHandler(
			client.UserServiceWrapper(service),
			// client.OrganizationServiceWrapper(service),
		),
		micro.WrapHandler(
			client.OrganizationServiceWrapper(service),
			client.MessageServiceWrapper(service),
		),
	)

	// Register Handler
	auth.RegisterAuthHandlerHandler(
		service.Server(),
		new(handler.AuthHandler),

		api.WithEndpoint(
			&api.Endpoint{
				Name:    "AuthHandler.Login",
				Path:    []string{"/v1/login"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "AuthHandler.Register",
				Path:    []string{"/v1/register"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
	)

	core.RegisterSMSHandlerHandler(
		service.Server(),
		new(handler.SMSHandler),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "SMSHandler.SendVerifyCode",
				Path:    []string{"/v1/sms/sendVerifyCode"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "SMSHandler.ValidateSimpleCode",
				Path:    []string{"/v1/sms/validateSimpleCode"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
	)

	// example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
