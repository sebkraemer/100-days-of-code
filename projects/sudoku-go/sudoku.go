// playground sudoku solver
package main

import (
	"fmt"
	"strings"
)

const N = 9

func main() {
	field := SudokuBoard{}

	ok := field.solveBt(0, 0)

	fmt.Printf("solved: %t\n%v\n", ok, field)
}

type SudokuBoard struct {
	data [N][N]int
}

// solve recursively with backtracking
func (sf *SudokuBoard) solveBt(col int, row int) bool {

	return true
}

func (sf *SudokuBoard) String() string {
	var b strings.Builder
	b.Grow(2 * N * N)
	for r := 0; r < N; r++ {
		fmt.Fprintf(&b, "%d %d %d %d %d %d %d %d %d\n", sf.data[r][0], sf.data[r][1], sf.data[r][2], sf.data[r][3], sf.data[r][4], sf.data[r][5], sf.data[r][6], sf.data[r][7], sf.data[r][8])
	}
	s := b.String()
	s = s[:len(s)-1] // cut off last line break
	return s
}

// incomplete,
func (sf *SudokuBoard) verifyBoard() bool {
	checkSlice := func(s []int) bool {
		return false
	}
	var ok bool
	for r := 0; r < N; r++ {
		ok = checkSlice(sf.data[r][:])
		if !ok {
			return false
		}
	}
	for c := 0; c < 0; c++ {
		elements := []int{}
		for i := 0; i < N; i++ {
			elements = append(elements, sf.data[i][c])
		}
		ok = checkSlice(elements)
		if !ok {
			return false
		}
	}
	return ok
}
