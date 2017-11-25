package orderbookSampler

type Token struct {
	// figure out how to do this with encapsulation
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	bid    float64
	ask    float64
	last   float64
}

func (t Token) getName() string {
	return t.Name
}

func (t Token) setName(name string) {
	t.Name = name
}

func (t Token) updatePrice(bid float64, ask float64, last float64) {
	t.bid = bid
	t.last = last
	t.ask = ask
}