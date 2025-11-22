package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func LoadEnv() AppConfig {
	var C AppConfig
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not present", err)
	}
	C.Port = os.Getenv("PORT")
	if C.Port == "" {
		C.Port = "8080"
	}

	return C
}
