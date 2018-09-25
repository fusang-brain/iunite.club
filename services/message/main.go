package main

import (
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/services/message/client"
	"iunite.club/services/message/handler"
	"iunite.club/services/message/subscriber"
	"os"

	example "iunite.club/services/message/proto/example"
	sms "iunite.club/services/message/proto/sms"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	if dbName == "" {
		dbName = "unite"
	}

	if host == "" {
		host = "localhost:27017"
	}
	// New Service
	service := ironic.NewService(
		micro.Name("iunite.club.srv.message"),
		micro.Version("latest"),
		micro.WrapHandler(
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.Profile{},
						&models.ValidateCode{},
					)
					return nil
				},
				monger.DBName(dbName),
				monger.Hosts([]string{
					host,
				}),
			),
		),
	)

	// Initialise service
	service.Init(
		micro.WrapHandler(
			client.TestServiceWrapper(service),
		),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	sms.RegisterSMSHandler(service.Server(), new(handler.SMSHandler))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.message", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.message", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
