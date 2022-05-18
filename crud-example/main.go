package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	a := 2
	b := 2

	result := sum(a, b)
	fmt.Printf("Result is: %d\n", result)
}
