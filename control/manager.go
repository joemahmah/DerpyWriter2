package control

import "github.com/joemahmah/DerpyWriter2/dict"

type InputLogic func(string, *[]dict.Dictionary)
type OutputLogic func(int, *[]dict.Dictionary)string

type InputManager struct {
	Dict *[]dict.Dictionary
	Logic InputLogic
}

type OutputManager struct {
	Dict *[]dict.Dictionary
	Logic OutputLogic
}

type Manager struct {
	IM InputManager
	OM OutputManager
	Dict *[]dict.Dictionary
}

func (om *OutputManager) GenText(words int) string {
	return om.Logic(words, om.Dict)
}

func (im *InputManager) ParseInput(input string) {
	im.Logic(input, im.Dict)

	for i := 0; i < len(*im.Dict); i++ {
		(*im.Dict)[i].Calculate()
	}
}
