package main

import "fmt"

func main() {
	fmt.Println(Factorial(3))
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// Example n = 5
// 5 x 4 x 3 x 2 x 1