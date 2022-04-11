package dictionary

// 7esl.com has a comprehensive list of 5 letter words, all conveniently located
// inside <li> tags
type Esl struct {
	url string `default:"https://7esl.com/5-letter-words/"`
}

func (s *Esl) Url() string {
	return s.url
}

func (s *Esl) Selectors() []string {
	return []string{
		"li",
	}
}
