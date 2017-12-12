package sampler

import (
	"io/ioutil"
	"encoding/json"
)

type GDAX struct {
	Orderbook []*Orderbook
	Tokens []*Token
}

func ConfigGDAX(file string) *GDAX {
	cb := new(GDAX)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cb.Tokens)
	json.Unmarshal(configFile, &cb.Orderbook)
	cb.UpdateTokens()
	cb.UpdateOrderbook()
	return cb
}

func (cb *GDAX) UpdateTokens() {
	for _, tkn := range cb.Tokens {
		tkn.updatePrice()
	}
}

func (cb *GDAX) UpdateOrderbook() {
	for _, ob := range cb.Orderbook {
		ob.updateEntries()
	}
}