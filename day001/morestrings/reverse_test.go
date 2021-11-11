package morestrings

import "testing"

func TestReverseString(t *testing.T) {
	// notice anonymous struct here
	cases := []struct {
		sample   string
		expected string
	}{
		{"foo", "oof"},
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseString(c.sample)
		if got != c.expected {
			t.Errorf("ReverseString(%s) (%s) != %s", c.sample, got, c.expected)
		}
	}

}
