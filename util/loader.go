package util

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, error) {

	if err := godotenv.Load(); err != nil {
		return "", errors.New("fail to load .env")
	}
	dbUri := os.Getenv("MONGODB_URI")
	if dbUri == "" {
		return "", errors.New("failed to load mongo db uri")
	}
	return dbUri, nil
}
