package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var denominator int = 2
	var numerator int = 8
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Println(err.Error())

	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)

	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}
func printMe(printValue string) {
	fmt.Println(printValue)
}

// must specify type of value being returned here
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
