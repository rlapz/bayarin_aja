package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvFrom(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error while loading `.env` file")
	}

	return os.Getenv(key)
}
