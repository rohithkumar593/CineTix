package dataAccessLayer

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDB struct {
	db *gorm.DB
}

func (postgresqlConnector *postgresDB) ConnectDB(connStr string) error {
	postgres_db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Error while connecting to DB", err)
		return err
	}
	postgresDBClient = postgres_db
	postgresqlConnector.db = postgresDBClient
	return nil
}

func (postgresqlConnector *postgresDB) CloseConnection() {
	postgresqlConnector.CloseConnection()
}

func GetPostgresClient() *gorm.DB {
	return postgresDBClient
}
