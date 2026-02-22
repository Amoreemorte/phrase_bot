package services

import "strings"

type Phraser struct {
	codePhrase string
}

func NewPhraser(codePhrase string) *Phraser {
	return &Phraser{
		codePhrase: codePhrase,
	}
}

func (p *Phraser) IsPhrasesMatch(matchingPhrase string) bool {
	return strings.EqualFold(p.codePhrase, matchingPhrase)
}
