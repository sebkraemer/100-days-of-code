package main

import "testing"

func TestString(t *testing.T) {
	sb := SudokuBoard{
		data: [N][N]int{{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{3, 4, 5, 6, 7, 8, 9, 0, 1},
			{5, 6, 7, 8, 9, 0, 1, 2, 3},
			{7, 8, 9, 0, 1, 2, 3, 4, 5},
			{9, 0, 1, 2, 3, 4, 5, 6, 7},
			{2, 3, 4, 5, 6, 7, 8, 9, 0},
			{4, 5, 6, 7, 8, 9, 0, 1, 2},
			{6, 7, 8, 9, 0, 1, 2, 3, 4},
			{8, 9, 0, 1, 2, 3, 4, 5, 6}},
	}
	var expected string = `1 2 3 4 5 6 7 8 9
3 4 5 6 7 8 9 0 1
5 6 7 8 9 0 1 2 3
7 8 9 0 1 2 3 4 5
9 0 1 2 3 4 5 6 7
2 3 4 5 6 7 8 9 0
4 5 6 7 8 9 0 1 2
6 7 8 9 0 1 2 3 4
8 9 0 1 2 3 4 5 69`

	// note: direct string conversion does not work
	got := sb.String()
	if got != expected {
		t.Errorf("wrong board representation; expected:\n%s\ngot:\n%s\n", got, expected)
	}

}
