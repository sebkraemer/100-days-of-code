package binarygap

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// gets be bits in reversed order but that's no problem for the task at hand
func getBinaryRepresentation(n int) []uint8 {
	res := []uint8{}
	for {
		res = append(res, uint8(n&1))
		n >>= 1 // n will be zero once all bits were processed
		if n == 0 {
			fmt.Println(res)
			break
		}
	}
	return res
}

func Solution(N int) int {
	var lastStart = -1
	var lastEnd = -1
	var count int

	bits := getBinaryRepresentation(N)
	if len(bits) == 0 {
		return 0
	}
	if bits[0] == 1 {
		lastStart = 0
	}
	for i := 1; i < len(bits); i++ {
		if bits[i] == 1 {
			if bits[i-1] == 0 && lastStart != -1 {
				lastEnd = i
				newCount := lastEnd - lastStart - 1
				if newCount > count {
					count = newCount
				}
				lastStart = -1
				continue
			}
		} else if bits[i] == 0 {
			if bits[i-1] == 1 {
				lastStart = i - 1
			}
			continue
		} else {
			panic("unexpected bit value!")
		}
	}
	return count
}
