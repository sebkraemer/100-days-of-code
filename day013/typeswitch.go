package main

import "fmt"

type withName interface {
	Greet() string
}

type person struct {
	name string
}

func (p person) Greet() string {
	return "I am " + p.name
}

func do(i interface{}) {
	switch v := i.(type) {
	// either persion or withName will match, apparently the first one wins

	case person:
		// type conversion
		fmt.Printf("it's a person: %s\n", person(v).name)

	//case person:
	//    fmt.Printf("it's a person: %s\n", v.(person).name)
	// 	compile error: ./prog.go:26:32: invalid type assertion: v.(person) (non-interface type person on left)
	case withName:
		// type assertion
		v = v.(person) // shadowing, it's said to me idiomatic to shadow the switched variabled
		fmt.Println(v.Greet())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	p := person{"max"}
	do(p)
}
