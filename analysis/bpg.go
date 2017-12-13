package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	_ "github.com/sjwhitworth/golearn/evaluation"
	_ "github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	rawData, err := base.ParseCSVToInstances("12-Dec-17.csv", false)
	if err != nil {
		panic(err)
	}

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

}