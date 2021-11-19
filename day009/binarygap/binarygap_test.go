package binarygap

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		sample   int
		expected int
	}{
		{0b0, 0},
		{0b01, 0},
		{0b1001, 2},
		{529, 4},
		{20, 1},
		{15, 0},
		{6, 0},
		{328, 2},
		{32, 0},
		{33, 4},
		{74901729, 4},
		{1376796946, 5},
		{66561, 9},
		{1162, 3},
	}

	for _, c := range cases {
		fmt.Printf("testing %d, expect %d gap size\n", c.sample, c.expected)
		got := Solution(c.sample)
		if got != c.expected {
			t.Errorf("binarygap.Solution(%d)=%d != %d", c.sample, got, c.expected)
		}
	}

}
