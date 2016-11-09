package main

import (
	"fmt"
	"github.com/joemahmah/DerpyWriter2/dict"
	"github.com/joemahmah/DerpyWriter2/control"
	"github.com/joemahmah/DerpyWriter2/output"
	"github.com/joemahmah/DerpyWriter2/inputFormat"
	"github.com/joemahmah/DerpyWriter2/tokenize"
	"flag"
	"io/ioutil"
)

func main() {

	lenVal := flag.Int("len",100,"The output length. Varies based on output logic.")
	outputLogic := flag.Int("ol",2,"Selects which logic to use when generating output. See documentation for list.")
	n := flag.Int("n",1,"IO depth.")

	inputType := flag.Int("it",2,"Determine where input is read from (0 stdin, 1 arg, 2 file).")

	flag.Parse()

	args := flag.Args()


	var toParse string

	switch *inputType {
		case 0:
			fmt.Println("Not yet implemented!")
			return
		case 1:
			toParse = args[0]
		case 2:
			contents, err := ioutil.ReadFile(args[0])
			if err != nil {
				panic(err)
			}
			toParse = string(contents)
		default:
			fmt.Println("Invalid input type!")
			return
	}

	ifm := control.InputFormatManager{Logic: inputFormat.SimpleTokenizeSpacePunctuation}
	toParse = ifm.Format(toParse)

	d := dict.NewDictionary(*n)
	d.SetTokenGen(dict.WordGenerator)

	dicts := make([]dict.Dictionary, 1)
	dicts[0] = *d

	tm := control.TokenizeManager{Dict: &dicts, Logic: tokenize.SimpleStringInput}
	tm.ParseInput(toParse)

	om := control.OutputManager{Dict: &dicts, Logic: output.SimpleMarkovOrderN(*n,*outputLogic)}

	fmt.Println(om.GenText(*lenVal))

}
