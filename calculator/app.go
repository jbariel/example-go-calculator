//
// Package main from a go CLI app.
//
// USAGE: <action> <init value> <action values...>
//  - <action> => one of (case insensitive) '+', '-', '*' '/', 'add', 'subtract', 'multiply', 'divide'
//  - <init value> => starting value (e.g. '1', '5.0252')
//  - <action values...> => one or more values to perform <action> on <init value>
//
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SayHi("User - I am your calculator!"))

	action, nums, err := checkArgs(os.Args)
	if nil == err {
		fmt.Printf("Will perform: '%s' on '%s'...\n", action, numsToString(nums))
		result, err := DoCalculate(action, nums)
		if nil == err {
			fmt.Printf("Value: %f\n\n", result)
		}
	}
	if nil != err {
		fmt.Println(err)
	}
}

func numsToString(nums []float32) string {
	strs := make([]string, 0, len(nums))
	for _, f := range nums {
		strs = append(strs, fmt.Sprintf("%f", f))
	}
	return strings.Join(strs, ", ")
}

func checkArgs(args []string) (action string, nums []float32, err error) {
	fmt.Printf("[DEBUG] Checking paramters: '%s'\n", strings.Join(args, ", "))
	if 4 > len(args) {
		err = errors.New("Usage: '" + args[0] + " action, init, vals...'\n\t=> must provide 2+ vals\n\n")
	}
	if nil == err {
		action = args[1]
		var num float64
		for _, val := range args[2:] {
			num, err = strconv.ParseFloat(val, 32)
			if nil == err {
				nums = append(nums, float32(num))
			}
		}
	}
	return
}
