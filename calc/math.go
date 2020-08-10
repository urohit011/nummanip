package calc

// Add is sued to add two numbers
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum
}
