package dictionary

import (
	"fmt"
	"strings"

	"github.com/DanielOchoa/wordle/scraper"
)

// REPL/cli tool to run scraper
// how?
// i := importer.New(Esl{})
// i.Scrape()
// i.Words()

type Importer struct {
	scraper *scraper.Scraper
	site    ScrapingSite
	words   []string
}

func New(site ScrapingSite) *Importer {
	return &Importer{scraper: scraper.New(site.Url()), site: site}
}

func (importer *Importer) Scrape() {
	fmt.Println("Importing...")
	fmt.Println("=====================================================")

	var wordList []string

	for _, selector := range importer.site.Selectors() {
		importer.scraper.OnElementFound(selector, func(e *scraper.Element) {
			wordList = append(wordList, strings.TrimSpace(e.Text))
		})
	}

	importer.words = wordList
}

// TODO:
// handle element found. In our case, insert into sqlite.
// validate it is a valid word.
