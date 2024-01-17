package configs

import (
	"log"
	"sync"

	"github.com/lpernett/godotenv"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Database struct {
		ConnectionString string `envconfig:"CONNECTION_STRING"`
	}
	Booking struct {
		Initialised int `envconfig:"INITIALISED"`
		Confirmed   int `envconfig:"CONFIRMED"`
		HoldTime    int `envconfig:"HOLD_TIME"`
	}
}

var AppConfig *Config = new(Config)
var configSetup sync.Once

func InitAppConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error encountered loading environment variables %v", err)
		return err
	}
	envconfig.MustProcess("", AppConfig)
	return nil
}
