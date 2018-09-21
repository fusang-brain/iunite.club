package main

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/srv/organization-srv/handler"
	clubPB "iunite.club/srv/organization-srv/proto/club"
	deptPB "iunite.club/srv/organization-srv/proto/department"
	jobPB "iunite.club/srv/organization-srv/proto/job"
	schoolPB "iunite.club/srv/organization-srv/proto/school"

	"iunite.club/srv/organization-srv/services"
	"os"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kit.iron.srv.organization-srv"),
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
		&models.User{},
		&models.Profile{},
		&models.Organization{},
		&models.UserClubProfile{},
		&models.School{},
		&models.OrganizationJob{},
		&models.JoinOrganizationAccept{},
	)

	ass := assistant.NewAssistant(
		assistant.Name("kit.iron.srv.origanization-srv"),
		assistant.Connection(connection),
		assistant.RegisterHandler(
			&handler.SchoolHandler{},
			&handler.ClubHandler{},
			&handler.JobHandler{},
			&handler.DepartmentHandler{},
		),
		assistant.RegisterService(
			&services.SchoolService{},
			&services.ClubService{},
			&services.DepartmentService{},
			&services.JobService{},
		),
	)

	// Initialise service
	service.Init()

	// Register Handler
	schoolPB.RegisterSchoolSrvHandler(
		service.Server(),
		ass.Handler("handler.SchoolHandler").(*handler.SchoolHandler),
	)

	clubPB.RegisterClubHandler(
		service.Server(),
		ass.Handler("handler.ClubHandler").(*handler.ClubHandler),
	)

	jobPB.RegisterJobHandler(
		service.Server(),
		ass.Handler("handler.JobHandler").(*handler.JobHandler),
	)

	deptPB.RegisterDepartmentHandler(
		service.Server(),
		ass.Handler("handler.DepartmentHandler").(*handler.DepartmentHandler),
	)
	// Register Struct as Subscriber
	// micro.RegisterSubscriber("kit.iron.srv.srv.organization-srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("kit.iron.srv.srv.organization-srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
