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
8 9 0 1 2 3 4 5 6`

	// note: direct string conversion does not work
	got := sb.String()
	if got != expected {
		t.Errorf("wrong board representation; expected:\n%s\ngot:\n%s\n", got, expected)
	}

}

func Test_solve(t *testing.T) {
	field := SudokuBoard{
		data: [N][N]int{{3, 0, 6, 5, 0, 8, 4, 0, 0},
			{5, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 8, 7, 0, 0, 0, 0, 3, 1},
			{0, 0, 3, 0, 1, 0, 0, 8, 0},
			{9, 0, 0, 8, 6, 3, 0, 0, 5},
			{0, 5, 0, 0, 9, 0, 6, 0, 0},
			{1, 3, 0, 0, 0, 0, 2, 5, 0},
			{0, 0, 0, 0, 0, 0, 0, 7, 4},
			{0, 0, 5, 2, 0, 6, 3, 0, 0}},
	}

	solved := field.solve()

	if !solved {
		t.Error("solve() returned false")
	}

	expected := `3 1 6 5 7 8 4 9 2
5 2 9 1 3 4 7 6 8
4 8 7 6 2 9 5 3 1
2 6 3 4 1 5 9 8 7
9 7 4 8 6 3 1 2 5
8 5 1 7 9 2 6 4 3
1 3 8 9 4 7 2 5 6
6 9 2 3 5 1 8 7 4
7 4 5 2 8 6 3 1 9`

	got := field.String()
	if field.String() != expected {
		t.Errorf("solved board unexpected, expected\n%s\ngot\n%s", expected, got)
	}
}
