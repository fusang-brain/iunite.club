package main

import (
	"os"
	"time"

	"iunite.club/services/core/proto/contacts"

	announce "iunite.club/services/core/proto/announce"
	"iunite.club/services/core/proto/conversation"
	recruitment "iunite.club/services/core/proto/recruitment"

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
		// micro.WrapHandler
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(rl.NewBucket(time.Second, 50), true),
			ironic.MongerWrapper(
				func(conn monger.Connection) error {
					conn.BatchRegister(
						&models.Approved{},
						&models.ApprovedFlow{},
						&models.User{},
						&models.RecruitmentForm{},
						&models.RecruitmentFormRecord{},
						&models.RecruitmentRecord{},
						&models.Announce{},
						&models.Conversation{},
						&models.ConversationNotice{},
						&models.SocialApplication{},
						&models.UserFriend{},
						// &models.User{},
						// &models.Profile{},
						&models.Role{},
						&models.RoleGroup{},
						&models.Permission{},
						&models.CasbinRule{},
						&models.Organization{},
						&models.UserClubProfile{},
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
	recruitment.RegisterRecruitmentHandler(service.Server(), new(handler.Recruitment))
	announce.RegisterAnnounceHandler(service.Server(), new(handler.Announce))
	iunite_club_srv_core_conversation.RegisterConversationHandler(service.Server(), new(handler.ConversationHandler))
	iunite_club_srv_core_contacts.RegisterContactsHandler(service.Server(), new(handler.ContactsHandler))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iunite.club.srv.core", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
