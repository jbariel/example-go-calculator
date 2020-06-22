package calculator

import (
	"errors"
	"fmt"
	"strings"
)

// HiString constant to say hi
const HiString string = "Hi "

// SayHi returns HiString + s
func SayHi(s string) string {
	return HiString + s
}

// DoCalculate takes an action (string) and an array of float32.  The first value
// in the array is the init value, the rest perform <action> on the init value
func DoCalculate(action string, nums []float32) (result float32, err error) {
	var fxn func(float32) float32

	switch strings.ToLower(action) {
	case "+", "add":
		fmt.Println("Attempting to add...")
		fxn = func(i float32) float32 { return result + i }
	case "-", "subtract":
		fmt.Println("Attempting to subtract...")
		fxn = func(i float32) float32 { return result - i }
	case "*", "multiply":
		fmt.Println("Attempting to multiply...")
		fxn = func(i float32) float32 { return result * i }
	case "/", "divide":
		fmt.Println("Attempting to divide...")
		fxn = func(i float32) float32 { return result / i }
	default:
		err = errors.New("I do not understand action '" + action + "'!!")
	}
	if nil == err {
		result = nums[0]
		for _, i := range nums[1:] {
			result = fxn(i)
		}
	}
	return
}
