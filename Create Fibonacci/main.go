package main

import "fmt"

func main() {
	FibonacciNumbers := CreateFibonacciNumbers(10)	// Create 10 element
	fmt.Println(FibonacciNumbers)
}

func CreateFibonacciNumbers(K int) []int {
	Fib := []int{}

	for i := 0; i < K; i++ {
		if i == 0 || i == 1 {
			Fib = append(Fib, i)
			continue
		}

		num := Fib[i-1] + Fib[i-2]
		Fib = append(Fib, num)
	}
	
	return Fib
}