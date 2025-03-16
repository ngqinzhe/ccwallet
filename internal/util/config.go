package util

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PostgreSqlCredentials PostgreSqlCredentials `yaml:"postgreSqlCredentials"`
}

type PostgreSqlCredentials struct {
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func InitConfig() *Config {
	data, err := os.ReadFile("./internal/config/debug.yml")
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML data: %v", err)
	}

	return &config
}
