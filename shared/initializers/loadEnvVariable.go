package initializers

import (
	"github.com/sirupsen/logrus"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("./shared/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logrus.Info("Env variables loaded successfully!")
}
