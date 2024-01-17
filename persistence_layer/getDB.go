package dataAccessLayer

import (
	"cine-tickets/configs"
	dbInterface "cine-tickets/interfaces/dbInterfaces"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

var (
	dbClient         map[string]dbInterface.DatabaseAccessObject = make(map[string]dbInterface.DatabaseAccessObject)
	initDbOnce       sync.Once
	postgresDBClient *gorm.DB
)

func initDBClients() {
	fmt.Println("starting to access")
	postgresConnector := postgresDB{}
	fmt.Println("done to access", postgresConnector)
	dbClient["postgres"] = &postgresConnector
}

func GetDbByName(name string) error {
	initDBClients()
	var err error
	initDbOnce.Do(func() {
		err = dbClient[name].ConnectDB(configs.AppConfig.Database.ConnectionString)
	})
	return err
}
