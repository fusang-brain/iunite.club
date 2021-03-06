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
	"iunite.club/services/organization/handler"
	"iunite.club/services/organization/subscriber"

	club "iunite.club/services/organization/proto/club"
	department "iunite.club/services/organization/proto/department"
	example "iunite.club/services/organization/proto/example"
	job "iunite.club/services/organization/proto/job"
	school "iunite.club/services/organization/proto/school"
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
		micro.Name("iunite.club.srv.organization"),
		micro.Version("latest"),
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(rl.NewBucket(time.Second, 50), true),
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.Profile{},
						&models.Organization{},
						&models.UserClubProfile{},
						&models.School{},
						&models.OrganizationJob{},
						&models.OrganizationAccept{},
					)

					return nil
				},
				monger.DBName(dbName),
				monger.Hosts([]string{
					host,
				}),
			),
			// wrappers.AuthWrapper(ser)
		),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	club.RegisterClubHandler(service.Server(), new(handler.ClubHandler))
	department.RegisterDepartmentHandler(service.Server(), new(handler.DepartmentHandler))
	job.RegisterJobHandler(service.Server(), new(handler.JobHandler))
	school.RegisterSchoolSrvHandler(service.Server(), new(handler.SchoolHandler))
	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.organization", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.organization", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
