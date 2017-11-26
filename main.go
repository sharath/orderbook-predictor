package main
import (
	"orderbook-predictor/sampler"
	"fmt"
)

func main() {
	gemini := new(sampler.Exchange)
	gemini.SetName("Gemini")
	gemini.UpdateTokens()
	fmt.Println(gemini.GetName())
	for i := 0; i < len(gemini.GetTokens()); i++ {
		fmt.Println(gemini.GetTokens()[i].Name + " " + fmt.Sprint(gemini.GetTokens()[i].Last))
	}
}