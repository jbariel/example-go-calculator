package calculator

import (
	"testing"
)

func TestSayHi(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"foo", "Hi foo"},
		{"bar", "Hi bar"},
		{"this is a LONG string", "Hi this is a LONG string"},
	}
	for _, c := range cases {
		got := SayHi(c.in)
		if c.expected != got {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}
