package main

import "fmt"

func status(x, y []int) {
	fmt.Printf("x: %v\ny: %v\n\n", x, y)
}

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	status(x, y)

	y[0] = 99 // will change x as well
	status(x, y)

	y = append(y, 55) // append() will also change x for as long as it's within x's len()
	status(x, y)

	y = append(y, 44)
	status(x, y)

	y = append(y, 10, 20, 30, 40) // length of x will not change
	status(x, y)

	y[0] = 0
	status(x, y) // copy! x[0] unchanged!
}

// output

// x: [1 2 3 4]
// y: [1 2]
//
// x: [99 2 3 4]
// y: [99 2]
//
// x: [99 2 55 4]
// y: [99 2 55]
//
// x: [99 2 55 44]
// y: [99 2 55 44]
//
// x: [99 2 55 44]
// y: [99 2 55 44 10 20 30 40]
//
// x: [99 2 55 44]
// y: [0 2 55 44 10 20 30 40]
