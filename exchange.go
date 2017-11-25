package orderbookSampler

import (
	"net/http"
	"encoding/json"
)

type Exchange struct {
	name string
	tickerurl string
	symbolsurl string
	tokens []Token
	gsymbols []GeminiSymbol
}

type GeminiSymbol struct {
	Name string `json:""`
}

func (e Exchange) Name() string {
	return e.name
}

func (e Exchange) setName(name string) {
	e.name = name
}

func (e Exchange) setTicker(tickerurl string, symbolurl string) bool {
	if e.name == "Gemini" {
		e.symbolsurl = symbolurl
		e.tickerurl = tickerurl
		response,err := http.Get(e.symbolsurl)
		if err != nil {
			return false
		}
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&e.gsymbols)
		e.tokens = []Token{*new(Token), *new(Token)} // this is bad code but it works
		bitcoin := e.tokens[0]
		ethereum := e.tokens[1]
	}


	return true
}