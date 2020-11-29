package sms

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

const TwilioAddress = "https://api.twilio.com/2010-04-01/Accounts/"

// SendText publishes a SMS message to a target number using one of the twilio's registered
// phone numbers.
func SendText(targetNo, message string) error {
	cfg, err := parseTwilioCredentials()
	if err != nil {
		log.Errorln(err)
		return err
	}

	data := url.Values{}
	data.Set("Body", message)
	data.Set("From", "+442393162353")
	data.Set("To", targetNo)

	client := &http.Client{}
	r, err := http.NewRequest(
		"POST",
		TwilioAddress+cfg.AccountSid+"/Messages.json",
		strings.NewReader(data.Encode()))

	response, err := client.Do(r)
	if err != nil {
		log.Errorln("Failed to send SMS", err)
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Debugln("Response from twilio " + string(body))

	return nil
}
