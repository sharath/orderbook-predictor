package sampler

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
)

type Token struct {
	// figure out how to do this with encapsulation
	Name   string  `json:"name"`
	Ticker string  `json:"ticker"`
	Bid    float64 `json:"bid,string"`
	Ask    float64 `json:"ask,string"`
	Last   float64 `json:"last,string"`
}

func (t *Token) GetName() string {
	return t.Name
}

func (t *Token) UpdatePrice() {
	response,err := http.Get(t.Ticker)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
