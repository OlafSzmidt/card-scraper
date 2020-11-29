package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/OlafSzmidt/card-scraper/internal/config"
	"github.com/OlafSzmidt/card-scraper/internal/scraper"
)

func main() {
	log.Info("Starting card-scraper...")

	cfg, err := config.ReadConfig()
	if err != nil {
		return
	}

	var scrapers []*scraper.Scraper

	for _, target := range cfg.Targets {
		log.WithFields(log.Fields{
			"url":   target.URL,
			"limit": target.Limit,
		}).Infoln("Creating a new scraper")

		sc, err := scraper.NewScraper(
			target.URL,
			make([]string, 0),
			target.Limit,
			target.HTMLSelector,
			cfg.Sms.Enabled,
		)

		if err != nil {
			return
		}

		scrapers = append(scrapers, sc)
	}

	var wg sync.WaitGroup

	for _, scraper := range scrapers {
		wg.Add(1)
		go scraper.Run(&wg)
	}

	wg.Wait()
}
