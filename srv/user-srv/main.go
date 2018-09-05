package main

import (
	"fmt"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/iron-kit/unite-services/secruity/wrappers"
	"github.com/iron-kit/unite-services/user-srv/handler"
	"github.com/iron-kit/unite-services/user-srv/models"
	"github.com/iron-kit/unite-services/user-srv/services"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"os"

	user "github.com/iron-kit/unite-services/user-srv/proto/user"
)

func main() {
	// New Service
	serviceName := "kit.iron.srv.user"
	service := micro.NewService(
		micro.Name("kit.iron.srv.user"),
		micro.Version("latest"),

		// 认证白名单
		micro.WrapHandler(wrappers.GenerateAuthWrapper(
			wrappers.NewWhiteItem(serviceName, "UserSrv.SigninByMobile"),
			wrappers.NewWhiteItem(serviceName, "UserSrv.ResetPasswordByMobile"),
			wrappers.NewWhiteItem(serviceName, "UserSrv.RegisterUserByMobile"),
		)),
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

	assistant := assistant.NewAssistant(
		assistant.Name("kit.iron.srv.user"),
		assistant.Connection(connection),
		assistant.RegisterHandler(&handler.User{}),
		assistant.RegisterService(&services.UserService{}),
	)

	// Initialise service
	service.Init()

	// Register Handler
	uh := assistant.Handler("handler.User").(*handler.User)
	user.RegisterUserSrvHandler(service.Server(), uh)
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// // Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.Example))

	// // Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
