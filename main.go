package main

import (
	"fmt"
	"time"
	slr "github.com/sharath/orderbook-predictor/sampler"
)

func sample(cb *slr.GDAX) {
	for {
		time.Sleep(2 * time.Second)
		n := time.Now()
		cb.UpdateTokens()
		cb.UpdateOrderbook()
		for i := 0; i < 2; i++ {
			fmt.Print(n.Format("02-Jan-06, 15:04PM") + ", ")
			fmt.Print(cb.Tokens[i].Name + ", ")
			fmt.Println(cb.Tokens[i].Price)
			fmt.Println("Bid:")
			for i := 0; i < len(cb.Orderbook[i].Bids); i++ {
				fmt.Print("Price:")
				fmt.Println(cb.Orderbook[i].Bids[i][0])
				fmt.Print("Amount:")
				fmt.Println(cb.Orderbook[i].Bids[i][1])
			}
			fmt.Println("Ask:")
			for i := 0; i < len(cb.Orderbook[i].Bids); i++ {
				fmt.Print("Price:")
				fmt.Println(cb.Orderbook[i].Bids[i][0])
				fmt.Print("Amount:")
				fmt.Println(cb.Orderbook[i].Bids[i][1])
			}
		}
	}
}

func main() {
	cb := slr.ConfigGDAX("config.json")
	go sample(cb)
	select {}
}