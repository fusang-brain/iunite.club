package main

import (
	"os"
	"time"

	rl "github.com/juju/ratelimit"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/services/core/handler"
	"iunite.club/services/core/subscriber"

	approved "iunite.club/services/core/proto/approved"
	example "iunite.club/services/core/proto/example"
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
		micro.Name("iunite.club.srv.core"),
		micro.Version("latest"),
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(rl.NewBucket(time.Second, 50), true),
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.Approved{},
						&models.ApprovedFlow{},
						&models.User{},
						// &models.User{},
						// &models.Profile{},
						// &models.Organization{},
						// &models.UserClubProfile{},
						// &models.School{},
						// &models.OrganizationJob{},
						// &models.OrganizationAccept{},
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
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	approved.RegisterApprovedHandler(
		service.Server(),
		new(handler.ApprovedHandler),
	)
	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
