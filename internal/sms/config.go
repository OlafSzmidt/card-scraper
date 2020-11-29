package sms

import (
	"errors"

	"github.com/kelseyhightower/envconfig"

	log "github.com/sirupsen/logrus"
)

type TwilioCredentials struct {
	AccountSid string `envconfig:"ACCOUNTSID"`
	AuthToken  string `envconfig:"AUTHTOKEN"`
}

func init() {
	parseTwilioCredentials()
}

func parseTwilioCredentials() (TwilioCredentials, error) {
	var cfg TwilioCredentials
	if err := envconfig.Process("", &cfg); err != nil {
		log.Errorln("Failed to read twilio credentials")
		return cfg, err
	}

	if len(cfg.AccountSid) <= 0 {
		return cfg, errors.New("Account SID has not been set")
	}

	if len(cfg.AuthToken) <= 0 {
		return cfg, errors.New("Account Auth Token has not been set")
	}

	return cfg, nil
}
