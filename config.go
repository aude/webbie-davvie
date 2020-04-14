package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	Dir  string
}

func NewConfig() (*Config, error) {
	config := new(Config)

	envConfig, err := parseEnv()
	if err != nil {
		return envConfig, fmt.Errorf("could not parse env variables: %s", err)
	}

	config = envConfig

	return config, nil
}

func parseEnv() (*Config, error) {
	config := new(Config)
	var value string
	var success bool

	if value, success = os.LookupEnv("PORT"); !success {
		return config, fmt.Errorf("%s is not set", "PORT")
	}
	config.Port = value
	if value, success = os.LookupEnv("DIR"); !success {
		return config, fmt.Errorf("%s is not set", "DIR")
	}
	config.Dir = value

	return config, nil
}
