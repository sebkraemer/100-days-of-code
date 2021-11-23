# Log of day 13

Don't know what got me there but I started playing around with the type switch. Applied my new knowledge about interface type, which I didn't have too look up 😃.

The outcome was quite interesting:

- "discovered" a faint of the thing between type assertion and type cast -- I really want to read up on it, maybe here: https://www.sohamkamani.com/golang/type-assertions-vs-type-conversions/ and spec here https://go.dev/ref/spec#Type_assertions
  - type switch on type -> type cast possible
  - type switch on interface type -> type assertion possible
- Apparently, the first mathing case wins (what else, really, if both match equally well?)

Also I was bold enough to file an errata on the *Learning Go* book, p. 151: the `i := i.(type)` should read ` i - i.(type)` to allow the shadowing. Let's see what comes back from O'Reilly.

See also [this day's short log entry](../log.md#day-13).

# Program snippet and output:

```Go
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
```

```text
% go run typeswitch.go
I don't know about type int!
I don't know about type string!
I don't know about type bool!
it's a person. person max
```
