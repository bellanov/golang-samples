package calculator

import "fmt"

// Calculator provides basic arithmetic operations.
// This demonstrates a simple service pattern in Go.
type Calculator struct{}

// NewCalculator creates a new Calculator instance.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add returns the sum of two integers.
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers.
// It returns an error if the divisor is zero.
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
