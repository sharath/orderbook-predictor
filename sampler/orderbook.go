package sampler

import (
	"net/http"
	"encoding/json"
)

type Orderbook struct {
	Name     string          `json:"name"`
	Book     string          `json:"book"`
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
}

func (ob *Orderbook) updateEntries() {
	response,_ := http.Get(ob.Book)
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ob)
}