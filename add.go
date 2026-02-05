// Package add provides a simple function to add two integers.
package add

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two integers.
// It takes two int parameters and returns their sum as an int.
func Add[T Number](a T, b T) T {
	return a + b
}
