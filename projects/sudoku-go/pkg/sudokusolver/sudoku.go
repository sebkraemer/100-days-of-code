// playground sudoku solver library
package sudokusolver

import (
	"fmt"
	"strings"
)

const N = 9

type SudokuBoard struct {
	Data [N][N]int `json:"data"`
}

// solve recursively with backtracking
//
// Find row, col of an unassigned cell
// If there is none, return true
//
// For digits from 1 to 9
//   if there is no conflict for digit at row,col
//     assign digit to row,col and recursively try fill in rest of grid
//     if recursion successful, return true
//     if !successful, remove digit and try another
// if all digits have been tried and nothing worked, return false to trigger backtracking
func (sf *SudokuBoard) Solve() bool {
	r, c, avail := sf.getNextAvailable()
	if !avail {
		return true
	}

	// TODO: would be interesting to try a heuristic
	for v := 1; v <= N; v++ {
		rowOk := sf.rowIsSafe(r, v)
		colOk := sf.colIsSafe(c, v)
		blockOk := sf.blockIsSafe(r, c, v)

		if rowOk && colOk && blockOk {
			sf.Data[r][c] = v
			if sf.Solve() {
				return true
			}
			sf.Data[r][c] = 0 // reset failed value
		}
	}
	return false
}

func (sf *SudokuBoard) rowIsSafe(row, value int) bool {
	for col := 0; col < N; col++ {
		if sf.Data[row][col] == value {
			return false
		}
	}
	return true
}

func (sf *SudokuBoard) colIsSafe(col, value int) bool {
	for row := 0; row < N; row++ {
		if sf.Data[row][col] == value {
			return false
		}
	}
	return true
}

func (sf *SudokuBoard) blockIsSafe(row, col, value int) bool {
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for ir := 0; ir < 3; ir++ {
		for ic := 0; ic < 3; ic++ {
			if sf.Data[startRow+ir][startCol+ic] == value {
				return false
			}
		}
	}
	return true
}

func (sf *SudokuBoard) getNextAvailable() (int, int, bool) {
	// todo: I wonder if we really have to start at zero each time, as long as we traverse left-to-right, top-to-bottom 🧐

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if sf.Data[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func (sf *SudokuBoard) String() string {
	var b strings.Builder
	b.Grow(2 * N * N)
	for r := 0; r < N; r++ {
		fmt.Fprintf(&b, "%d %d %d %d %d %d %d %d %d\n", sf.Data[r][0], sf.Data[r][1], sf.Data[r][2], sf.Data[r][3], sf.Data[r][4], sf.Data[r][5], sf.Data[r][6], sf.Data[r][7], sf.Data[r][8])
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
		ok = checkSlice(sf.Data[r][:])
		if !ok {
			return false
		}
	}
	for c := 0; c < 0; c++ {
		elements := []int{}
		for i := 0; i < N; i++ {
			elements = append(elements, sf.Data[i][c])
		}
		ok = checkSlice(elements)
		if !ok {
			return false
		}
	}
	return ok
}
