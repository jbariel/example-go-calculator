package calculator

import (
	"fmt"
	"strings"
)

// SayHi returns "Hi " + s
func SayHi(s string) string {
	return "Hi " + s
}

// DoCalculate takes an action (string) and an array of float32.  The first value
// in the array is the init value, the rest perform <action> on the init value
func DoCalculate(action string, nums []float32) (result float32, error string) {
	result = nums[0]
	switch strings.ToLower(action) {
	case "+", "add":
		fmt.Println("Attempting to add...")
		for _, i := range nums[1:] {
			result = result + i
		}
	case "-", "subtract":
		fmt.Println("Attempting to subtract...")
		for _, i := range nums[1:] {
			result = result - i
		}
	case "*", "multiply":
		fmt.Println("Attempting to multiply...")
		for _, i := range nums[1:] {
			result = result * i
		}
	case "/", "divide":
		fmt.Println("Attempting to divide...")
		for _, i := range nums[1:] {
			result = result / i
		}
	default:
		error = "I do not understand action '" + action + "'!!"
	}
	return
}
