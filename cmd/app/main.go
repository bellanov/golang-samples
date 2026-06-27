package main

import (
	"fmt"

	"github.com/bellanov/golang-samples/internal/calculator"
)

func main() {
	calc := calculator.NewCalculator()

	result := calc.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	result = calc.Subtract(10, 4)
	fmt.Printf("10 - 4 = %d\n", result)

	result = calc.Multiply(6, 7)
	fmt.Printf("6 * 7 = %d\n", result)

	result, err := calc.Divide(20, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("20 / 4 = %d\n", result)
	}

	_, err = calc.Divide(20, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
