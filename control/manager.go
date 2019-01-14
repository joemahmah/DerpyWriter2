package control

import (
	"github.com/joemahmah/DerpyWriter2/dict"
	"github.com/joemahmah/DerpyWriter2/output"
)

type InputFormatLogic func(string) string
type TokenizeLogic func(string, *[]dict.Dictionary)
type OutputLogic func(int, *[]dict.Dictionary) []output.Sentence

type InputFormatManager struct {
	Logic InputFormatLogic
}

type TokenizeManager struct {
	Dict *[]dict.Dictionary
	Logic TokenizeLogic
}

type OutputManager struct {
	Dict *[]dict.Dictionary
	Logic OutputLogic
}

type Manager struct {
	IFM InputFormatManager
	TM TokenizeManager
	OM OutputManager
	Dict *[]dict.Dictionary
}

func (om *OutputManager) GenText(words int) []output.Sentence {
	return om.Logic(words, om.Dict)
}

func (tm *TokenizeManager) ParseInput(input string) {
	tm.Logic(input, tm.Dict)

	for i := 0; i < len(*tm.Dict); i++ {
		(*tm.Dict)[i].Calculate()
	}
}

func (ifm *InputFormatManager) Format(input string) string{
	return ifm.Logic(input)
}
