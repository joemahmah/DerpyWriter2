package dict

type Token interface {
	GetName() string
	GetRarity() int
	AsString() string
	IncrementRarity()
	SetName(string)
	SetStrVal(string)
	GetTokensBefore(int) map[string]int
	GetTokensAfter(int) map[string]int
	GetMaxDepth() int
	SetMaxDepth(int)
	GetTokensBeforeWeighted() [][]string
	GetTokensAfterWeighted() [][]string
	Calculate()
}

type Word struct {
	Rarity int
	Name string
	StrVal string
	TokensBefore []map[string]int
	TokensAfter []map[string]int
	MaxDepth int
	TokensAfterWeighted [][]string
	TokensBeforeWeighted [][]string
}

func (w Word) GetMaxDepth() int {
	return w.MaxDepth
}

func (w Word) GetName() string{
	return w.Name
}

func (w Word) GetRarity() int{
	return w.Rarity
}

func (w Word) AsString() string{
	return w.StrVal
}

func (w *Word) SetMaxDepth(MaxDepth int) {
	w.MaxDepth = MaxDepth
}

func (w *Word) SetStrVal(strVal string) {
	w.StrVal = strVal
}

func (w *Word) SetName(name string) {
	w.Name = name
}

func (w *Word) IncrementRarity() {
	w.Rarity += 1
}

func (w *Word) GetTokensBefore(index int) map[string]int {
	return w.TokensBefore[index]
}


func (w *Word) GetTokensAfter(index int) map[string]int {
	return w.TokensAfter[index]
}

func (w *Word) GetTokensAfterWeighted() [][]string {
	return w.TokensAfterWeighted
}

func (w *Word) GetTokensBeforeWeighted() [][]string {
	return w.TokensBeforeWeighted
}

func (w *Word) Calculate(){
	w.TokensAfterWeighted = make([][]string, 0)
	w.TokensBeforeWeighted = make([][]string, 0)

	for _, val := range  w.TokensAfter{
		tmpSlice := make([]string,0)

		for str, mapval := range val {
			for i := 0; i < mapval; i++{
				tmpSlice = append(tmpSlice, str)
			}
		}

		w.TokensAfterWeighted = append(w.TokensAfterWeighted, tmpSlice)
	}

	for _, val := range w.TokensBefore{
		tmpSlice := make([]string,0)


		for str, mapval := range val {
			for i := 0; i < mapval; i++{
				tmpSlice = append(tmpSlice, str)
			}
		}

		w.TokensBeforeWeighted = append(w.TokensBeforeWeighted, tmpSlice)
	}
}

func WordGenerator(name string, maxDepth int) Token {
	word := &Word{Name: name, StrVal: name, Rarity: 1, MaxDepth: maxDepth}
	word.TokensBefore = make([]map[string]int, maxDepth)
	word.TokensAfter = make([]map[string]int, maxDepth)

	for i := 0; i < maxDepth; i++ {
		word.TokensBefore[i] = make(map[string]int)
		word.TokensAfter[i] = make(map[string]int)
	}

	//Token is an interface; as such, the type Token cannot exist. Tokens are references to actual implementations, hence &Type rather than Type
	return word
}


