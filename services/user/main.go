package main

import (
	"time"

	"github.com/iron-kit/go-ironic"
	rl "github.com/juju/ratelimit"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"iunite.club/services/user/client"

	// "github.com/iron-kit/go-ironic/wrappers"
	"github.com/iron-kit/monger"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"iunite.club/models"

	// "iunite.club/services/user/client"
	"os"

	"iunite.club/services/user/handler"

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
			ratelimit.NewHandlerWrapper(rl.NewBucket(time.Second, 50), true),
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.User{},
						&models.Profile{},
						&models.School{},
						&models.UserClubProfile{},
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

	// Register Handler
	user.RegisterUserSrvHandler(service.Server(), new(handler.UserSrv))
	secruity.RegisterSecruityHandler(service.Server(), new(handler.Secruity))

	// Initialise service
	service.Init(
		micro.WrapHandler(
			client.MessageServiceWrapper(service),
		),
	)
	// service.Init(
	// 	micro.WrapHandler(
	// 		client.MessageServiceWrapper(service),
	// 		wrappers.AuthWrapper(
	// 			service,
	// 			wrappers.NewWhiteItem("UserSrv.SigninByMobile"),
	// 			wrappers.NewWhiteItem("UserSrv.ResetPasswordByMobile"),
	// 			wrappers.NewWhiteItem("UserSrv.RegisterUserByMobile"),
	// 		),
	// 	),
	// )

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("iunite.club.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
