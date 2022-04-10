package dictionary

import "github.com/gocolly/colly/v2"

// Scraper for 5 words from https://7esl.com/5-letter-words/ to start with.
// Scrape all <li> elements.
// Perhaps this is a wrapper for an individual site?

type Element = colly.HTMLElement

type Scraper struct {
	Words  []string
	client *colly.Collector
	*Options
}

type Options struct {
	url      string
	selector *colly.HTMLElement
	callback func(e *Element)
}

func New(url string, opts *Options) *Scraper {
	scraper := &Scraper{client: colly.NewCollector(), Options: opts}
	return scraper
}

func (s *Scraper) Scrape(selector string, callback func(e *Element)) {
	s.client.OnHTML(selector, callback)
	s.client.Visit(s.Options.url)
}
