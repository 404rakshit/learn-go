package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("New lesson")

	numerator := 11
	denominator := 3
	result, remainder, err := intDivision(numerator, denominator)

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the interger division is %v", result)
	default:
		fmt.Printf("The result of the interger division is %v with remainder %v", result, remainder)
	}

}
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New(`cannot dicide by zero`)
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
