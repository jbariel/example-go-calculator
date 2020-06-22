package main

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

func TestDoCalculate(t *testing.T) {
	for _, c := range []struct {
		action   string
		nums     []float32
		expected float32
	}{
		{"add", []float32{0, 1, 2}, 3},
		{"ADD", []float32{0, 1, 2}, 3},
		{"AdD", []float32{0, 1, 2}, 3},
		{"+", []float32{0, 1, 2}, 3},
		{"+", []float32{0, 2, 4}, 6},
		{"+", []float32{0, -1, 2}, 1},
		{"subtract", []float32{0, 1, 2}, -3},
		{"SUBTRACT", []float32{0, 1, 2}, -3},
		{"SUBtraCT", []float32{0, 1, 2}, -3},
		{"-", []float32{0, 1, 2}, -3},
		{"-", []float32{0, 2, 4}, -6},
		{"-", []float32{0, -1, 2}, -1},
		{"multiply", []float32{1, 1, 2}, 2},
		{"MULTIPLY", []float32{1, 1, 2}, 2},
		{"MULTIply", []float32{1, 1, 2}, 2},
		{"*", []float32{0, 1, 2}, 0},
		{"*", []float32{1, 2, 4}, 8},
		{"*", []float32{1, -1, 2}, -2},
		{"divide", []float32{2, 1, 2}, 1},
		{"DIVIDE", []float32{2, 1, 2}, 1},
		{"DIVide", []float32{2, 1, 2}, 1},
		{"/", []float32{0, 1, 2}, 0},
		{"/", []float32{8, 2, 4}, 1},
		{"/", []float32{2, -1, 2}, -1},
	} {
		got, err := DoCalculate(c.action, c.nums)
		if nil != err {
			t.Error(err)
		}
		if c.expected != got {
			t.Errorf("DoCalculate(%q, %q) == %f, want %f", c.action, numsToString(c.nums), got, c.expected)
		}
	}
}
