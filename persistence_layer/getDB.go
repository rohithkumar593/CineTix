package dataAccessLayer

import (
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
	err := dbClient[name].ConnectDB("user=kamasala.r dbname=kamasala.r sslmode=disable")
	return err
}
