package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port           int    `yaml:"port"`
	Env            string `yaml:"env"`
	DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
}

func NewConfig(configPath string) *Config {
	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
