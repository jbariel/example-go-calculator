package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World")

	args := os.Args
	if 4 > len(args) {
		fmt.Printf("Usage: '%s action, val...'\n\t=> must provide 2+ val\n\n", args[0])
	} else {
		action := os.Args[1]
		numStr := os.Args[2:]
		nums := make([]float32, 0, len(numStr))
		for _, strVal := range numStr {
			numVal, err := strconv.ParseFloat(strVal, 32)
			if nil != err {
				fmt.Printf("Could not parse value of '%s'!\n", strVal)
			} else {
				nums = append(nums, float32(numVal))
			}
		}

		fmt.Printf("Will perform: '%s' on \n", action)
		fmt.Println("\t", nums)

		res := nums[0]
		switch strings.ToLower(action) {
		case "+", "add":
			fmt.Println("Attempting to add...")
			for _, i := range nums[1:] {
				res = res + i
			}
		case "-", "subtract":
			fmt.Println("Attempting to subtract...")
			for _, i := range nums[1:] {
				res = res - i
			}
		case "*", "multiply":
			fmt.Println("Attempting to multiply...")
			for _, i := range nums[1:] {
				res = res * i
			}
		case "/", "divide":
			fmt.Println("Attempting to divide...")
			for _, i := range nums[1:] {
				res = res / i
			}
		default:
			fmt.Printf("I do not understand action '%s'!!\n\n", action)
		}

		fmt.Println("Value: ", res)
	}
}
