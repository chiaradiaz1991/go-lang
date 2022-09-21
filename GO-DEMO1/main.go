package main

import (
	"fmt"

	mascotpckg "example.com/go-demo-1/mascot-pckg"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascotpckg.BestMascot())
	fmt.Println(quote.Go())
}
