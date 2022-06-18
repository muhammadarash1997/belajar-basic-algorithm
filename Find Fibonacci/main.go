package main

import "fmt"

func main() {
	FibNumber := Fibonacci(10)	// Find Nth element
	fmt.Println(FibNumber)
}

func Fibonacci(K int) int {
	if K == 0 || K == 1 {
		return K
	}

	return Fibonacci(K-1) + Fibonacci(K-2)
}