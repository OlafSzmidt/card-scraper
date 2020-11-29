package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// AppConfiguration is the initial yaml configuration of the application,
// defining all endpoints.
type AppConfiguration struct {
	Targets []struct {
		URL          string `yaml:"url"`
		Limit        int    `yaml:"limit"`
		HTMLSelector string `yaml:"selector"`
	}
	Sms struct {
		Enabled bool `yaml:"enabled"`
	}
}

// ReadConfig will load the config.yml file.
func ReadConfig() (*AppConfiguration, error) {
	log.Infoln("Reading config.yml")

	f, err := os.Open("config.yml")
	if err != nil {
		log.Errorln("Failed to open file", err)
		return nil, err
	}
	defer f.Close()

	var cfg AppConfiguration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Errorln("Failed to decode the configuration", err)
		return nil, err
	}

	log.Infoln(cfg)
	return &cfg, nil
}
