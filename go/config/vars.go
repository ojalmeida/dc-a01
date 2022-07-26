package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var Config Configuration

func init() {

	rawData, err := os.ReadFile(configurationFile)

	if err != nil {
		panic(errors.New("Unable to read configuration file at " + configurationFile + ": " + err.Error()))
	}

	err = yaml.Unmarshal(rawData, &Config)

	if err != nil {
		panic(errors.New("Unable to read configuration file at " + configurationFile + ": " + err.Error()))
	}

}

var configurationFile = "/etc/config.yaml"
