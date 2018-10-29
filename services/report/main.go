package main

import (
	"os"
	"iunite.club/models"
	"github.com/iron-kit/monger"
	"github.com/iron-kit/go-ironic"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/services/report/handler"
	// "iunite.club/services/report/subscriber"

	pb "iunite.club/services/report/proto"
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
		micro.Name("iunite.club.srv.report"),
		micro.Version("latest"),
		micro.WrapHandler(
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.Report{},
						&models.ReportTemplate{},
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
	pb.RegisterReportHandler(service.Server(), new(handler.Report))
	// Register Struct as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.report", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.report", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
