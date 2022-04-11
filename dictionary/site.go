package dictionary

// All sites have the same interface
type ScrapingSite interface {
	Url() string
	Selectors() []string
}
