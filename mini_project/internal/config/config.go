package config

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/wiryawanyurih/ecommerce-api/vendor/github.com/joho/godotenv"
)

type Config struct {
	Listen struct {
		Host string
		Port string
	}
	DB struct {
		Database string
		URI      string
		Username string
		Password string
	}
	Redis struct {
		URI string
	}
	Stripe struct {
		Key           string
		WebhookUrl    string
		WebhookSecret string
	}
	Test struct {
		Database   string
		DBURI      string
		DBUsername string
		DBPassword string
	}
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}

		instance.Listen.Host = os.Getenv("HOST")
		instance.Listen.Port = os.Getenv("PORT")

		instance.DB.Username = os.Getenv("DB_USERNAME")
		instance.DB.Password = os.Getenv("DB_PASSWORD")
		instance.DB.Database = os.Getenv("DB_NAME")
		instance.DB.URI = os.Getenv("DB_URI")

		instance.Redis.URI = os.Getenv("REDIS_URI")
	})
	return instance
}
