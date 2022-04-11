package scraper

import "github.com/gocolly/colly/v2"

// Steps:
// Scraper for 5 words from https://7esl.com/5-letter-words/ to start with.
// Scrape all <li> elements.

// Wrapper for colly scraper

type Element = colly.HTMLElement

type Scraper struct {
	client *colly.Collector
	url    string
}

func New(url string) *Scraper {
	return &Scraper{client: colly.NewCollector(), url: url}
}

func (s *Scraper) OnElementFound(selector string, callback func(e *Element)) {
	s.client.OnHTML(selector, callback)
}

func (s *Scraper) Visit() {
	s.client.Visit(s.url)
}
