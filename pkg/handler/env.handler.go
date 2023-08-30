package handler

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvHandler() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	return
}
