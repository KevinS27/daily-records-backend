package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var cfg Config

type Config struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	URL string `yaml:"url"`
}

func ReadConfig() *Config {
	configFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalln("failed to read config file")
		return nil
	}

	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Fatalln("failed to unmarshal config")
		return nil
	}

	return &cfg

}
