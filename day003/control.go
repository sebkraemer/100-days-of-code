package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//runAllIfStatements()
	fourForStatements()
}

// if statements
func runAllIfStatements() {
	ifWithExternalDeclaration()
	ifWithImplicitDeclaration()
}

func ifWithExternalDeclaration() {
	n := rand.Intn(20)
	if n == 3 {
		fmt.Println("you really hit three!")
	} else if n == 0 {
		fmt.Println("zarrruuuuuuu")
	} else {
		fmt.Println("boring") // this will always be hit because rand.Intn will deterministically return 1
	}
}

func ifWithImplicitDeclaration() {
	if n := rand.Intn(20); n != 1 {
		fmt.Println(n)
	}
}

func fourForStatements() {
	// 1. complete
	for i := 0; i < 7; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. condition-only
	i := 0
	for i < 7 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// 3. infinite loop, also uses break and continue
	i = -5
	for {
		if i < 0 {
			i++
			continue // don't print negative numbers
		}
		if i >= 7 {
			break // not an infinite loop after all
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// 4. range for loops

	// with index and value for slice
	list := []int{0, 1, 2, 3, 4, 5, 6}
	for i, v := range list {
		fmt.Print("i:", i, " v:", v, " ")
	}
	fmt.Println()

	// value only with ignored value (_); an unused variable would be a compile time error!
	for _, v := range list {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// key only with single v, skipping index
	for v := range list {
		fmt.Print(v, " ")
	}
	fmt.Println()

	m := map[string]int{"fizz": 3, "buzz": 5}
	for k, v := range m {
		fmt.Print("k:", k, " v:", v, " ")
	}
	fmt.Println()
	for k := range m {
		fmt.Print(k, " ") // this will not always be "fizz buzz" :o
	}
	fmt.Println()

	s := "hello 🌞 ✌️" // hello sunshine
	fmt.Println(s)
	for i, r := range s {
		if i == 0 {
			fmt.Printf("type of i: %T\n", i)
			fmt.Printf("type of r: %T\n", r) // int32 aka rune, i.e. not byte
		}
		fmt.Print("i:", i, " r:", r, " ")
		// i:0 r:104 i:1 r:101 i:2 r:108 i:3 r:108 i:4 r:111 i:5 r:32 i:6 r:127774 i:10 r:32 i:11 r:9996 i:14 r:65039
		// notice the skips in the index after multi-byte strings were processed
	}

	fmt.Println()
	for _, c := range s {
		switch even := c % 2; c {
		case 'l':
			fmt.Print(" found l;")
			fmt.Print("even?: ", even == 0)
		default:
			fmt.Print(c, "even?: ", even == 0)
		}

	}
	fmt.Println()

}
