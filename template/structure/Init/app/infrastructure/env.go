package infrastructure

import (
	"gogengotest/app/interfaces"
	"github.com/joho/godotenv"
)

// Load is load configs from a env file.
func Load(logger interfaces.Logger) {
	err := godotenv.Load()
	if err != nil {
		logger.LogError("%s", err)
	}
}