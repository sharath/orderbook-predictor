package main

import (
	slr "github.com/sharath/orderbook-predictor/sampler"
	"fmt"
	"time"
	"os"
	"path"
)

func main() {
	cb := slr.ConfigGDAX("config_minimal.json")
	os.MkdirAll("data", os.ModePerm)
	dataPath := path.Join("data", time.Now().Format("02-Jan-06.csv"))
	file, _ := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
	var index uint64 = 0
	for range time.NewTicker(time.Second).C {
		poll := func() {
			if file.Name() != path.Join("data", time.Now().Format("02-Jan-06.csv")) {
				file.Close()
				file, _ = os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
			}
			n := time.Now()
			cb.UpdateTokens()
			cb.UpdateOrderbook()
			for i := 0; i < len(cb.Tokens); i++ {
				for j := 0; j < len(cb.Orderbook[i].Bids); j++ {
					line := fmt.Sprintf("%d, %s, %s, Bid, %s, %s", index, n.Format("15, 04, 05.000"), cb.Tokens[i].Name, cb.Orderbook[i].Bids[j][0], cb.Orderbook[i].Bids[j][1])
					fmt.Println(line)
					fmt.Fprintln(file, line)
				}
				line := fmt.Sprintf("%d, %s, %s, Market, %f", index, n.Format("15, 04, 05.000"), cb.Tokens[i].Name, cb.Tokens[i].Price)
				fmt.Println(line)
				fmt.Fprintln(file, line)
				for j := 0; j < len(cb.Orderbook[i].Asks); j++ {
					line := fmt.Sprintf("%d, %s, %s, Ask, %s, %s", index, n.Format("15, 04, 05.000"), cb.Tokens[i].Name, cb.Orderbook[i].Asks[j][0], cb.Orderbook[i].Asks[j][1])
					fmt.Println(line)
					fmt.Fprintln(file, line)
				}
			}
		}
		go poll()
		index++
	}
	select {}
}