package main

import "fmt"

func main() {
	fmt.Println(Power(2, 2))
}

func Power(n int, p int) int {
	if p == 1 {
		return n
	}

	return n * Power(n, p-1)
}

// Example n = 5 & p = 3
// 5 x 5 x 5