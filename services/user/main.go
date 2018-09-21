package main

import (
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/wrappers"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"
	"iunite.club/services/user/handler"
	"os"

	user "iunite.club/services/user/proto"
	secruity "iunite.club/services/user/proto/secruity"
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
		micro.Name("iunite.club.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.Profile{},
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
	service.Init(
		micro.WrapHandler(
			wrappers.AuthWrapper(
				service,
				wrappers.NewWhiteItem("UserSrv.SigninByMobile"),
				wrappers.NewWhiteItem("UserSrv.ResetPasswordByMobile"),
				wrappers.NewWhiteItem("UserSrv.RegisterUserByMobile"),
			),
		),
	)

	// Register Handler
	user.RegisterUserSrvHandler(service.Server(), new(handler.UserSrv))
	secruity.RegisterSecruityHandler(service.Server(), new(handler.Secruity))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
