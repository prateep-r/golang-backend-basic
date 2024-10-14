package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Server        Server
	AccessControl AccessControl
	Database      Database
	Header        Header
}

type Server struct {
	Hostname string `env:"HOSTNAME"`
	Port     string `env:"PORT,notEmpty"`
}

type AccessControl struct {
	AllowOrigin string `env:"ACCESS_CONTROL_ALLOW_ORIGIN"`
}

type Database struct {
	PostgresURL string `env:"POSTGRES_URL"` // example: postgres://postgres:password@localhost:5432/dbname?sslmode=disable
	RedisURL    string `env:"REDIS_URL"`
}

type Header struct {
	RefIDHeaderKey string `env:"REF_ID_HEADER_KEY,notEmpty"`
}

var once sync.Once
var config Config

func prefix(e string) string {
	if e == "" {
		return ""
	}

	return fmt.Sprintf("%s_", e)
}

func C(envPrefix string) Config {
	once.Do(func() {
		opts := env.Options{
			Prefix: prefix(envPrefix),
		}

		var err error
		config, err = parseEnv[Config](opts)
		if err != nil {
			log.Fatal(err)
		}
	})

	return config
}
