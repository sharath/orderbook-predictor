package orderbookSampler

import (
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
)

type Exchange struct {
	name   string
	tokens []Token
}

func (e *Exchange) GetName() string {
	return e.name
}

func (e *Exchange) SetName(name string) {
	e.name = name
	if name == "Gemini" {
		config, err := ioutil.ReadFile("orderbookSampler/api-data/gemini.json") // TODO clean up relative path
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		json.Unmarshal(config, &e.tokens)
	}
}

func (e *Exchange) UpdateTokens() {
	for i := 0;i < len(e.tokens);i++ {
		e.tokens[i].UpdatePrice()
	}
}

func (e *Exchange) GetTokens() []Token {
	return e.tokens
}