package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

//Getenv for get environtment variable
func Getenv(keys string) string {
	if os.Getenv(keys) == "" {
		return ""
	}
	return os.Getenv(keys)
}
