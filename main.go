package main

import (
	"fmt"
	"time"
	slr "github.com/sharath/orderbook-predictor/sampler"
	"os"
	"path"
)

func sample(cb *slr.GDAX) {
	dataPath := path.Join("data", time.Now().Format("02-Jan-06.csv"))
	file, _ := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
	for {
		if file.Name() != path.Join("data", time.Now().Format("02-Jan-06.csv")) {
			file.Close()
			file, _ = os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
		}
		time.Sleep(time.Second*5)
		// GDAX keeps banning my ip
		n := time.Now()
		cb.UpdateTokens()
		cb.UpdateOrderbook()

		for i := 0; i < len(cb.Tokens); i++ {
			fmt.Fprintln(file,n.Format("02-Jan-06, 15:04:05.000PM"))
			fmt.Fprint(file,cb.Tokens[i].Name + ", Market, ")
			fmt.Fprintln(file,cb.Tokens[i].Price)
			for j := 0; j < len(cb.Orderbook[i].Bids); j++ {
				fmt.Fprint(file,cb.Tokens[i].Name + ", Bid, ")
				fmt.Fprint(file,cb.Orderbook[i].Bids[j][0])
				fmt.Fprint(file,", ")
				fmt.Fprintln(file,cb.Orderbook[i].Bids[j][1])
			}
			for j := 0; j < len(cb.Orderbook[i].Asks); j++ {
				fmt.Fprint(file,cb.Tokens[i].Name + ", Ask, ")
				fmt.Fprint(file,cb.Orderbook[i].Asks[j][0])
				fmt.Fprint(file,", ")
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