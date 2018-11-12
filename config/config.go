package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Can Not Load ENV Files: " + err.Error())
	}
}
