package main

import (
	"os"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/services/storage/handler"
	proto "iunite.club/services/storage/proto"
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
		micro.Name("iunite.club.srv.storage"),
		micro.Version("latest"),
		micro.WrapHandler(
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.File{},
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
	service.Init()

	// Register Handler
	proto.RegisterStorageHandler(
		service.Server(),
		new(handler.Storage),
	)

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.storage", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.storage", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
