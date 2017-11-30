package sampler

import (
	"net/http"
)

type Token struct {
	Name   string  `json:"name"`
	Ticker string  `json:"ticker"`
	Bid    float64 `json:"bid,string"`
	Ask    float64 `json:"ask,string"`
	Price  float64 `json:"price,string"`
}

func (t Token) getName() string {
	return t.Name
}

func (t Token) getPrice() (float64, float64, float64) {
	return t.Price, t.Bid, t.Ask
}

func (t *Token) updatePrice() {
	response,_ := http.Get(t.Ticker)
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&t)
}