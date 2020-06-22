package calculator

import (
	"testing"
)

func TestSayHi(t *testing.T) {
	for _, c := range []struct {
		in string
	}{
		{"foo"},
		{"bar"},
		{"this is a LONG string"},
	} {
		got := SayHi(c.in)
		if (HiString + c.in) != got {
			t.Errorf("SayHi(%q) == %q, want %q", c.in, got, HiString+c.in)
		}
	}
}
