package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

// Load is load configs from a env file.
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("%s", err)
	}
}