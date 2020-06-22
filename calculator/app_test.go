package main

import (
	"testing"
)

func eql(exp []float32, got []float32) (result bool) {
	result = len(exp) == len(got)
	if result {
		for i, v := range exp {
			result = result && v == got[i]
		}
	}
	return
}
func TestMain(t *testing.T) {
	main()
}

func TestCheckArgs(t *testing.T) {
	for _, c := range []struct {
		in     []string
		action string
		nums   []float32
	}{
		{[]string{"app_test.go", "+", "1.5", "1.5"}, "+", []float32{1.5, 1.5}},
		{[]string{"app_test.go", "ADD", "1.5", "1.5"}, "ADD", []float32{1.5, 1.5}},
		{[]string{"app_test.go", "-", "1.5", "1.5", "1.5"}, "-", []float32{1.5, 1.5, 1.5}},
		{[]string{"app_test.go", "*", "1.5", "1.5"}, "*", []float32{1.5, 1.5}},
		{[]string{"app_test.go", "/", "1.5", "1.5"}, "/", []float32{1.5, 1.5}},
	} {
		gotAction, gotNums, err := checkArgs(c.in)
		if nil != err {
			t.Error(err)
		} else if c.action != gotAction && !eql(c.nums, gotNums) {
			t.Errorf("checkArgs(%q) == (%q, %q), want (%q, %q)", c.in, gotAction, numsToString(gotNums), c.action, numsToString(c.nums))
		}
	}
}
