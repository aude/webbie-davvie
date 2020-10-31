package main

import (
	"fmt"
	"os"
	"regexp"
)

type Config struct {
	Port string
	Dir  string
}

func (config *Config) validate() error {
	// Port
	validPort := regexp.MustCompile(`^[0-9]+$`)
	if !validPort.MatchString(config.Port) {
		return fmt.Errorf("configured port \"%s\" is not valid", config.Port)
	}

	// Dir
	if stat, err := os.Stat(config.Dir); err != nil || !stat.IsDir() {
		return fmt.Errorf("configured dir \"%s\" is not valid", config.Dir)
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := new(Config)

	// take config from multiple sources
	//
	// priority, from least to most:
	//
	// - config file (not implemented)
	// - environment variable
	// - command line switch (not implemented)

	envConfig := NewConfigFromEnv()
	config.apply(envConfig)

	// always validate config
	if err := config.validate(); err != nil {
		return config, err
	}

	return config, nil
}

func (config *Config) apply(add *Config) {
	if add.Port != "" {
		config.Port = add.Port
	}
	if add.Dir != "" {
		config.Dir = add.Dir
	}
}

// create Config from environment variables
func NewConfigFromEnv() *Config {
	config := new(Config)

	config.Port = os.Getenv("PORT")
	config.Dir = os.Getenv("DIR")

	return config
}
