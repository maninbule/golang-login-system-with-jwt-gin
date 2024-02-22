package loadinit

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVaria() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
