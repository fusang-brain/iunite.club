package main

import (
	"os"
	"iunite.club/models"
	"github.com/iron-kit/monger"
	pb "iunite.club/services/approved/proto"
	"github.com/iron-kit/go-ironic"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/services/approved/handler"
	// example "iunite.club/services/approved/proto/example"
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
		micro.Name("iunite.club.srv.approved"),
		micro.Version("latest"),
		micro.WrapHandler(
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.ApprovedTemplate{},
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
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))
	pb.RegisterApprovedHandler(service.Server(), new(handler.Approved))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
