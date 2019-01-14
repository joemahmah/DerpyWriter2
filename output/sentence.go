package output

import "strings"

type Sentence struct {
	Words []string
	Punctuation string
}

func (s *Sentence) AddWord(word string) {
	s.Words = append(s.Words, word)
}

func (s *Sentence) SetPunctuation(punct string) {
	s.Punctuation = punct
}

func (s *Sentence) AsString() string {
	var output string = ""

	for _, word := range s.Words {
		output += " " + word
	}

	output += s.Punctuation

	return output
}

func IsPunctuation(word string) bool {
	formatted := strings.TrimSpace(word)

	switch formatted {
		case ".":
			fallthrough
		case "!":
			fallthrough
		case "?":
			return true
		default:
			return false
	}
}


func NewSentence() Sentence {
	var s Sentence

	s = Sentence{}
	s.Words = make([]string, 0)

	return s
}

