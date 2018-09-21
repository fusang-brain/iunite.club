package main

import (
	"fmt"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/srv/secruity/client"
	"iunite.club/srv/secruity/handler"
	"iunite.club/srv/secruity/services"
	"os"

	auth "iunite.club/srv/secruity/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.srv.secruity"),
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
	fmt.Println("db host:", host)
	connection, err := monger.Connect(
		monger.DBName(dbName),
		monger.Hosts([]string{
			host,
		}),

		// monger.PoolLimit(500)
	)

	if err != nil {
		panic(err.Error())
	}

	connection.BatchRegister(
		&models.User{},
		&models.Profile{},
	)

	ass := assistant.NewAssistant(
		assistant.Name("kit.iron.srv.secruity"),
		assistant.Connection(connection),
		assistant.RegisterHandler(&handler.Auth{}),
		assistant.RegisterService(&services.TokenService{}),
	)

	// Register Handler
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))
	auth.RegisterAuthHandler(service.Server(), ass.Handler("handler.Auth").(*handler.Auth))

	// Initialise service
	service.Init(
		micro.WrapHandler(
			client.UserWrapper(service),
		),
	)

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.secruity", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.secruity", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
