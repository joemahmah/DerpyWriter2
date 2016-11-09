package tokenize

import (
	"github.com/joemahmah/DerpyWriter2/dict"
	"strings"
)

func SimpleStringInput(input string, dictionary *[]dict.Dictionary) {

	words := strings.Fields(input)

	for _, word := range words {
		(*dictionary)[0].AddToken(word)
	}

}

