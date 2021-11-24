package main

import "fmt"

func main() {
	// declare variables
	var a, b, c int

	// assign values
	a = 10
	b = 4

	// arithmetic operations

	// addition of a and b
	c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)

	// subtraction of a and b
	c = a - b
	fmt.Printf("%d - %d: %d\n", a, b, c)

	// multiplication of a and b
	c = a * b
	fmt.Printf("%d * %d: %d\n", a, b, c)

	// division of a and b
	d := float32(a) / float32(b)
	fmt.Printf("%d + %d: %.2f\n", a, b, d)
}
