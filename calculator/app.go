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
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	SayHi("User - I am your calculator!")

	action, nums, err := checkArgs(os.Args)
	if nil == err {
		fmt.Printf("Will perform: '%s' on \n", action)
		fmt.Println("\t", nums)
		result, err := DoCalculate(action, nums)
		if nil == err {
			fmt.Println("Value: ", result)
		}
	}
	if nil != err {
		fmt.Println(err)
	}
}

func checkArgs(args []string) (action string, nums []float32, err error) {
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
