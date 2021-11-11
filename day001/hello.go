package main

import (
	"fmt"

	"github.com/sebkraemer/100-days-of-code/day001/morestrings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(quote.Go())
	fmt.Println(morestrings.ReverseString(quote.Go()))
}
