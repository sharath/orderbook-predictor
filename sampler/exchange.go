package sampler

import (
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
	"strings"
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
	config, err := ioutil.ReadFile("sampler/api-data/" + strings.ToLower(e.name) + ".json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	json.Unmarshal(config, &e.tokens)
}

func (e *Exchange) UpdateTokens() {
	for i := 0; i < len(e.tokens); i++ {
		e.tokens[i].UpdatePrice()
	}
}

func (e *Exchange) GetTokens() []Token {
	return e.tokens
}