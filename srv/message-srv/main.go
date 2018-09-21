package main

import (
	"iunite.club/srv/message-srv/handler"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"os"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.srv.srv.message-srv"),
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

		// monger.PoolLimit(500)
	)

	if err != nil {
		panic(err.Error())
	}

	connection.BatchRegister(
		&models.ValidateCode{},
	)

	ass := assistant.NewAssistant(
		assistant.Name("kit.iron.srv.srv.message-srv"),
		assistant.Connection(connection),
		assistant.RegisterHandler(
			handler.SMSHandler{}
		)
	)

	// Initialise service
	service.Init()

	// Register Handler

	// Register Struct as Subscriber

	// Register Function as Subscriber

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
