package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var Env string

func init() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Env = os.Getenv("ENV")
}

const (
	Dev  string = "DEV"
	Sit  string = "SIT"
	UAT  string = "UAT"
	Prod string = "PROD"
)

func IsDevEnv() bool {
	return strings.ToUpper(Env) == Dev
}

func IsSitEnv() bool {
	return strings.ToUpper(Env) == Sit
}

func IsUATEnv() bool {
	return strings.ToUpper(Env) == UAT
}

func IsProdEnv() bool {
	return strings.ToUpper(Env) == Prod
}
