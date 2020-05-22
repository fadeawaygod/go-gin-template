package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ = loadEnvFile()

/*
GetEnv obtain config from env variable, if not exist, load from env file.
if neither was found, return empty string
*/
func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func loadEnvFile() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("load .env file failed.")
	}
	return err
}
