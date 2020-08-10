package calc

import (
	"errors"
)

// Add is sued to add two numbers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Provide more than 2 arguments")
	}

	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}
