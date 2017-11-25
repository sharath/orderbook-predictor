package orderbookSampler

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
)

type Token struct {
	// figure out how to do this with encapsulation
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	bid    float64
	ask    float64
	last   float64
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
	json.NewDecoder(response.Body).Decode(&t)
}