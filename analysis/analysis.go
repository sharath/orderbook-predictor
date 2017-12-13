package main

import (
	"os"
	"time"
	"bufio"
	"strings"
	"strconv"
)

type MarketEntry struct {
	Name string
	Time time.Time
	Price float64
}

func ParsePrices(file *os.File) []MarketEntry {
	r := bufio.NewReader(file)
	var entries []MarketEntry
	for {
		in, _, err := r.ReadLine()
		line := string(in)
		split := strings.Split(line, ", ")

		if len(split) < 6 {
			break
		}
		if strings.Contains(split[5], "Market") {
			p, _ := strconv.ParseFloat(split[6], 2)
			t, _ := time.Parse("02-Jan-06.csv150405.000", os.Args[1]+split[1]+split[2]+split[3])
			entries = append(entries, MarketEntry{Name:split[4], Time:t, Price: p})
		}
		if err != nil {
			break
		}
	}
	return entries
}

func main() {
	dataFile, _ := os.Open(os.Args[1])
	defer dataFile.Close()
	entries := ParsePrices(dataFile)
	var prices []float64
	var times []time.Time
	for _, entry := range entries {
		prices = append(prices, entry.Price)
		times = append(times, entry.Time)
	}
}