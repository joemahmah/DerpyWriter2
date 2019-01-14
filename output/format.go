package output

func FormatPlain(sentences []Sentence) string {
	var output string = ""

	for _, sentence := range sentences {
		output += sentence.AsString()
	}

	return output
}

func FormatHTML(sentences []Sentence) string {
	output := FormatPlain(sentences)

	//TODO: convert to lines, add tags

	return output
}
