//
// Package calculator from a go CLI app.
//
// USAGE: <action> <init value> <action values...>
//  - <action> => one of (case insensitive) '+', '-', '*' '/', 'add', 'subtract', 'multiply', 'divide'
//  - <init value> => starting value (e.g. '1', '5.0252')
//  - <action values...> => one or more values to perform <action> on <init value>
//
package calculator

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World")
	SayHi("World")
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

		res, err := DoCalculate(action, nums)
		if "" == err {
			fmt.Println("Value: ", res)
		} else {
			fmt.Println(err)
		}
	}
}

func checkArgs() {

}
