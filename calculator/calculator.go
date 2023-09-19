// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("square root of a negative number is not allowed: %f", input)
	}
	return math.Sqrt(input), nil
}
