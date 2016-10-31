package dict

import (
	"errors"
	"math/rand"
)

type Dictionary struct {
	Tokens map[string]Token
	TokensNoMap []Token
	TokenGen func(string, int) Token
	lastWords []string
	MaxDepth int
}

func NewDictionary(maxDepth int) *Dictionary {
	return &Dictionary{Tokens: make(map[string]Token), lastWords: make([]string, maxDepth), MaxDepth: maxDepth}
}

func (d *Dictionary) GetTokenCount() int {
	return len(d.Tokens)
}

func (d *Dictionary) SetTokenGen(tokenGen func(string, int) Token) {
	d.TokenGen = tokenGen
}

func (d *Dictionary) AddSpecialToken(token string, tokenGenerator func(string, int) Token) {
	t, err := d.GetToken(token)

	if err == nil {
		t.IncrementRarity()
		d.fillBeforeAfterTokens(t)
		return
	}

	d.Tokens[token] = tokenGenerator(token, d.MaxDepth)
	d.fillBeforeAfterTokens(d.Tokens[token])
}

func (d *Dictionary) AddToken(token string) error {

	if d.TokenGen == nil {
		return errors.New("Token Generator is absant.")
	}

	t, err := d.GetToken(token)

	if err == nil {
		t.IncrementRarity()
		d.fillBeforeAfterTokens(t)
		return nil
	}

	d.Tokens[token] = d.TokenGen(token, d.MaxDepth)
	d.fillBeforeAfterTokens(d.Tokens[token])
	return nil
}


func (d *Dictionary) AddPrefabToken(token Token) {
	t, err := d.GetToken(token.GetName())

	if err == nil {
		t.IncrementRarity()
		d.fillBeforeAfterTokens(t)
		return
	}

	d.Tokens[token.GetName()] = token
	d.fillBeforeAfterTokens(token)
}

func (d *Dictionary) fillBeforeAfterTokens(token Token) {
	name := token.GetName()

	for i := d.MaxDepth; i > 0; i-- {
		tB,err := d.GetToken(d.lastWords[i-1])
		tA,err2 := d.GetToken(name)
		if err == nil {
			tB.GetTokensAfter(i-1)[name] = tB.GetTokensAfter(i-1)[name] + 1
		}

		if err2 == nil && d.lastWords[i-1] != "" {
			tA.GetTokensBefore(i-1)[d.lastWords[i-1]] = tA.GetTokensBefore(i-1)[d.lastWords[i-1]] + 1
		}
	}

	d.cycleLastWords(token)
}

func (d *Dictionary) cycleLastWords(token Token) {

	for i := d.MaxDepth; i > 1; i-- {
		d.lastWords[i-1] = d.lastWords[i-2]
	}

	d.lastWords[0] = token.GetName()

}

func (d *Dictionary) Contains(name string) bool {
	if d.Tokens[name] != nil{
		return true
	}

	return false
}

func (d *Dictionary) Calculate() {
	d.TokensNoMap = make([]Token, len(d.Tokens))

	idx := 0
	for _, token := range d.Tokens {
		token.Calculate()
		d.TokensNoMap[idx] = token
		idx++
	}

}

func (d *Dictionary) GetToken(name string) (Token, error) {
	token := d.Tokens[name]

	if token != nil {
		return token,nil
	}

	return nil,errors.New("Invalid key!")
}

func (d *Dictionary) GetRandomTokenUnweighted() Token {
	return d.TokensNoMap[rand.Intn(len(d.TokensNoMap))]
}
