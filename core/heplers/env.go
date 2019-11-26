package heplers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(e string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	return os.Getenv(e)
}
