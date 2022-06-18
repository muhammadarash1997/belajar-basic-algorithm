package main

import "fmt"

func main() {
	fmt.Println(SumSeries(1, 2, 5))
	fmt.Println(SumSeries(5, 1, 10))
}

func SumSeries(min int, gap int, max int) int {
	if min == max {
		return min
	}

	return min + SumSeries(min+gap, gap, max)
}

// Example min = 1, gap = 2, max = 5
// 1 + 3 + 5

// Example min = 5, gap = 1, max = 10
// 5 + 6 + 7 + 8 + 9 + 10