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
	fmt.Println(gemini.GetTokens())
}