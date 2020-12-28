package util

import (
	"log"
	"os"

	"git.woa.com/sr-devops/console-server/consts"
	"github.com/joho/godotenv"
)

// LoadEnv : Load environment variables in file .env
func LoadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = consts.DEFAULT_ENV
	}
	err := godotenv.Load("config/.env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
