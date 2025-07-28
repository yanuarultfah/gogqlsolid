package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"username"`
	Password     string `json:"password"`
	DBName       string `json:"database"`
	SSLMode      string `json:"sslmode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

func LoadConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()
	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return cfg
}
