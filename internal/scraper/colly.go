package scraper

import (
	"net/url"
	"sync"
	"time"

	"github.com/enescakir/emoji"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"github.com/OlafSzmidt/card-scraper/internal/sms"
)

// Scraper represents a single target scraping instance.
type Scraper struct {
	collector  *colly.Collector
	target     *url.URL
	limit      int
	smsEnabled bool
	selector   string
}

// NewScraper constructs a new Scraper struct, validating all data in process.
func NewScraper(target string, allowedDomains []string, limit int, selector string, smsEnabled bool) (*Scraper, error) {
	parsedTarget, err := url.Parse(target)
	if err != nil {
		log.WithFields(log.Fields{
			"targetURL": target,
		}).Errorln("Failed to parse target URL for the scraper")
		return nil, err
	}

	for _, domain := range allowedDomains {
		if _, err := url.Parse(domain); err != nil {
			log.WithFields(log.Fields{
				"allowedDomain": domain,
			}).Errorln("Failed to parse a domain URL for the scraper")
			return nil, err
		}
	}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		RandomDelay: time.Duration(limit) * time.Second,
	})

	return &Scraper{
		target:     parsedTarget,
		collector:  c,
		limit:      limit,
		selector:   selector,
		smsEnabled: smsEnabled,
	}, nil
}

// Run starts a Scraper goroutine at fixed limit intervals. Every Scraper.Limit seconds
// a new visit to the website will be triggered (with random deviance provided by colly).
func (s *Scraper) Run(waitGroup *sync.WaitGroup) {
	limiter := time.NewTicker(time.Duration(s.limit) * time.Second)

	for {
		<-limiter.C
		log.WithFields(log.Fields{
			"url": s.target.String(),
		}).Debugln("Limit time expired; invoking a scrape")
		s.scrape()
	}
}

func (s *Scraper) scrape() {
	s.collector.OnHTML(s.selector, func(e *colly.HTMLElement) {
		log.WithFields(log.Fields{
			"url": s.target.String(),
		}).Warnln(emoji.Warning + emoji.Warning + emoji.Warning + " Product is available!")

		if s.smsEnabled {
			sms.SendText("+4407857091682", "Hello olaf!")
		}
	})

	log.WithFields(log.Fields{
		"targetURL": s.target.String(),
	}).Debugln("Visiting a site.")

	s.collector.Visit(s.target.String())
}
