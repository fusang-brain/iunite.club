package main

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/micro/go-api"
	apiHandler "github.com/micro/go-api/handler/api"
	"github.com/micro/go-log"
	"iunite.club/api/navo/client"
	"iunite.club/api/navo/handler"
	auth "iunite.club/api/navo/proto/auth"
	"os"

	"github.com/micro/go-micro"
	// "iunite.club/api/navo/client"
	// "iunite.club/api/navo/handler"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.api.navo"),
		micro.Version("latest"),
	)

	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	if dbName == "" {
		dbName = "unite"
	}

	if host == "" {
		host = "localhost:27017"
	}

	connection, err := monger.Connect(
		monger.DBName(dbName),
		monger.Hosts([]string{
			host,
		}),
	)

	if err != nil {
		panic(err.Error())
	}

	ass := assistant.NewAssistant(
		assistant.Name("kit.iron.api.navo"),
		assistant.Connection(connection),
		assistant.RegisterHandler(
			&handler.AuthHandler{},
		),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		// micro.WrapHandler(client.ExampleWrapper(service)),
		micro.WrapHandler(client.SecruityWrapper(service)),
	)

	// Register Handler
	auth.RegisterAuthHandlerHandler(service.Server(), ass.Handler("handler.AuthHandler").(*handler.AuthHandler),
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
				Path:    []string{"/v1/login"},
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
