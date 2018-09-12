package utils

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"iunite.club/models"
)

func GetAssistant(srvs ...assistant.Servicer) *assistant.Assistant {
	connection, err := monger.Connect(
		monger.DBName("unite"),
		monger.Hosts([]string{
			// "localhost",
			"127.0.0.1:27017",
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

	assistant := assistant.NewAssistant(
		assistant.Connection(connection),
		assistant.RegisterService(srvs...),
	)

	return assistant
}
