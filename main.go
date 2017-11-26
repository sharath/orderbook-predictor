package main
import (
	obs "orderbookSampler/orderbookSampler"
	"fmt"
)

func main() {
	gemini := new(obs.Exchange)
	gemini.SetName("Gemini")
	gemini.UpdateTokens()
	fmt.Println(gemini.GetName())
	for i := 0; i < len(gemini.GetTokens()); i++ {
		fmt.Println(gemini.GetTokens()[i].Name + " " + fmt.Sprint(gemini.GetTokens()[i].Last))
	}
}