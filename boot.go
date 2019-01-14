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

	flag.Parse()

	args := flag.Args()

	inputFormatManager := control.InputFormatManager{Logic: inputFormat.SimpleTokenizeSpacePunctuation}

	baseDictionary := dict.NewDictionary(*n)
	baseDictionary.SetTokenGen(dict.WordGenerator)

	//DW2 is implemented to use multiple dictionaries
	//We need to assign the actual dictionary to the dictionary manager
	dictionaries := make([]dict.Dictionary, 1)
	dictionaries[0] = *baseDictionary

	tokenizeManager := control.TokenizeManager{Dict: &dictionaries, Logic: tokenize.SimpleStringInput}

	//Read through all the files
	for fileIdx := 0; fileIdx < len(args); fileIdx++ {
		contents, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		toParse := inputFormatManager.Format(string(contents))

		tokenizeManager.ParseInput(toParse)
	}

	//Create the output manager
	outputManager := control.OutputManager{Dict: &dictionaries, Logic: output.SimpleMarkovOrderN(*n,*outputLogic)}

	//Print output
	sentenceOutput := outputManager.GenText(*lenVal)
	fmt.Println(output.FormatPlain(sentenceOutput))

}
