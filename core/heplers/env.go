package heplers

import (
	"os"
)

func Getenv(e string) string {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env")
	// }

	return os.Getenv(e)
}
