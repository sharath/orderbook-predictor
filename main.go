package main

import (
	"fmt"
	"time"
	slr "github.com/sharath/orderbook-predictor/sampler"
	"os"
	"path"
)

func sample(cb *slr.GDAX) {
	file,_ := os.Create(path.Join("data", "file.txt"))
	defer file.Close()
	for {
		time.Sleep(0.5*time.Second)
		n := time.Now()
		cb.UpdateTokens()
		cb.UpdateOrderbook()

		for i := 0; i < len(cb.Tokens); i++ {
			fmt.Fprint(file,n.Format("02-Jan-06, 15:04PM, "))
			fmt.Fprint(file,cb.Tokens[i].Name + ", ")
			fmt.Fprintln(file,cb.Tokens[i].Price)
			for j := 0; j < len(cb.Orderbook[i].Bids); j++ {
				fmt.Fprint(file,cb.Tokens[i].Name + ", Bid, Price:")
				fmt.Fprint(file,cb.Orderbook[i].Bids[j][0])
				fmt.Fprint(file,", Amount:")
				fmt.Fprintln(file,cb.Orderbook[i].Bids[j][1])
			}
			for j := 0; j < len(cb.Orderbook[i].Asks); j++ {
				fmt.Fprint(file,cb.Tokens[i].Name + ", Ask, Price:")
				fmt.Fprint(file,cb.Orderbook[i].Asks[j][0])
				fmt.Fprint(file,", Amount:")
				fmt.Fprintln(file,cb.Orderbook[i].Asks[j][1])
			}
		}
	}
}

func main() {
	cb := slr.ConfigGDAX("config.json")
	os.MkdirAll("data", os.ModePerm)

	go sample(cb)
	select {}
}