package test_utils

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"github.com/iron-kit/unite-services/user-srv/models"
)

func GetAssistant(srvs ...assistant.Servicer) *assistant.Assistant {
	connection, err := monger.Connect(
		monger.DBName("unite"),
		monger.Hosts([]string{
			"localhost",
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
		assistant.Connection(connection),
		assistant.RegisterService(srvs...),
	)

	return assistant
}
