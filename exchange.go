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

func (e Exchange) Name() string {
	return e.name
}

func (e Exchange) setName(name string) {
	e.name = name
	if name == "Gemini" {
		config, err := ioutil.ReadFile("./api-data/gemini.json")
		if err != nil {
			fmt.Println("unable to read json configuration file")
			os.Exit(-1)
		}
		json.Unmarshal(config, &e.tokens)
	}
}