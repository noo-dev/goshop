package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	TEST_ENV       = "testing"
	PRODUCTION_ENV = "production"

	DATABASE_TIMEOUT     = 5 * time.Second
	PRODUCT_CACHING_TIME = 5 * time.Minute
)

type Schema struct {
	Environment string `env:"environment"`
	HttpPort    int    `env:"http_port"`
	DatabaseUri string `env:"database_uri"`
}

var cfg Schema

func init() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	environment := os.Getenv("environment")
	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil && environment != TEST_ENV {
		log.Fatalf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}
}

func GetConfig() *Schema {
	return &cfg
}
