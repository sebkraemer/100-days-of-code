package main

import (
	"fmt"

	"github.com/sebkraemer/100-days-of-code/projects/fints-go/pkg/parser"
)

func main() {

	p := parser.FinTSVisitor()
	fmt.Println("%T\n", p)
}
