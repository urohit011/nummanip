package calc

import (
	"errors"
)

// Add is sued to add two numbers
func Add(numbers ...int) int {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 arguments"), sum
	}

	for _, num := range numbers {
		sum += num
	}
	return sum
}
