package output

import (
	"bytes"
	"math/rand"
	"github.com/joemahmah/DerpyWriter2/dict"
)

func strSliceContains(slice []string, str string) bool {

	for _, elem := range slice {
		if elem == str {
			return true
		}
	}

	return false

}

//TODO: Write this function
//Runs in O(n^2) time
func markovSquareLikelyCanidates(n1canidates []string, canidates []string) []string {

	//copy markovRemoveUnlikelyCanidates, but rather than delete, append current element

	for idex := len(canidates)-1; idex >=0; idex-- {
		if strSliceContains(n1canidates, canidates[idex]){
			canidates = append(canidates,canidates[idex])
		}
	}

	return canidates
}

//Runs in O(n^2) time
func markovRemoveUnlikelyCanidates(n1canidates []string, canidates []string) []string {
	for idex := len(canidates)-1; idex >=0; idex-- {
		if !strSliceContains(n1canidates, canidates[idex]){
			canidates = append(canidates[:idex],canidates[idex+1:]...)
		}
	}

	return canidates
}

//Runs in O(1) time
func markovNoCanidateChange(n1canidates []string, canidates []string) []string {
	return canidates
}


//Note: Slow as fuck
func SimpleMarkovOrderN(n int, nOrderFunc int) func(int, *[]dict.Dictionary)string{

	var orderFunction func([]string,[]string)[]string

	//Used to determine which function should be used to gen text
	switch nOrderFunc {
		case 1:
			orderFunction = markovSquareLikelyCanidates
		case 2:
			orderFunction = markovRemoveUnlikelyCanidates
		default:
			orderFunction = markovNoCanidateChange
	}

	function := func(words int, dicts *[]dict.Dictionary) string{
		dictionary := (*dicts)[0]

		var output bytes.Buffer
		var lastTokens []dict.Token = make([]dict.Token, n)

		//Choose random first word
		current := dictionary.GetRandomTokenUnweighted()
		output.WriteString(current.AsString())
		output.WriteString(" ")

		lastTokens[0] = current

		var canidates []string
		var n1canidates []string

		for i := 1; i < words; i++ {
			canidates = make([]string, 0)
			n1canidates = make([]string, 0)

			//Choose canidates based on prior n tokens
			if lastTokens[0] != nil {
				canidates = append(canidates, lastTokens[0].GetTokensAfterWeighted()[0]...)
				n1canidates = append(n1canidates, lastTokens[0].GetTokensAfterWeighted()[0]...)
			}

			for currentN := 1; currentN < n; currentN++ {
				if lastTokens[currentN] != nil {
					canidates = append(canidates, orderFunction(n1canidates,lastTokens[currentN].GetTokensAfterWeighted()[currentN])...)
				}
			}

			//Choose new token from all possible canidates
			var err error
			choice := rand.Intn(len(canidates))
			current, err = dictionary.GetToken(canidates[choice])

			if err == nil {
				//Write to string
				output.WriteString(current.AsString())
				output.WriteString(" ")

				//Cycle last tokens
				for j := 0; j < n-1; j++ {
					lastTokens[j+1] = lastTokens[j]
				}
				lastTokens[0] = current
			} else {
				i--
			}
		}

		return output.String()
	}

	return function
}

//Only uses first dictionary; only uses 1 back for markov chain
func SimpleMarkov(words int, dicts *[]dict.Dictionary) string {
	return SimpleMarkovOrderN(1, 0)(words, dicts)
}

func SimpleRandom(words int, dicts *[]dict.Dictionary) string {
	dict := (*dicts)[0]

	var output bytes.Buffer

	for i := 0; i < words; i++ {
		output.WriteString(dict.GetRandomTokenUnweighted().AsString())
		output.WriteString(" ")
	}

	return output.String()
}

