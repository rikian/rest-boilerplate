package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// for more information, see --> https://github.com/joho/godotenv/issues/43#issuecomment-503183127
const projectDirName = "rest-api"

func LoadEnvFile() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error when load .env file\nerror: %v", err.Error())
	}
}
