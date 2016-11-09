package inputFormat

import "strings"

func SimpleTokenizeSpacePunctuation(input string) string {
	var output string
	output = strings.Replace(input, ".", " . ", -1)
	output = strings.Replace(output, ",", " , ", -1)
	output = strings.Replace(output, "!", " ! ", -1)
	output = strings.Replace(output, ";", " ; ", -1)
	output = strings.Replace(output, ":", " : ", -1)
	output = strings.Replace(output, "?", " ? ", -1)
	output = strings.Replace(output, "\t", " ", -1)
	output = strings.Replace(output, "\n", " ", -1)
	output = strings.Replace(output, "  ", " ", -1)

	return output
}
