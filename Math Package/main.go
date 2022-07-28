package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Min(1, 2))  // 1
	fmt.Println(math.Min(-2, 2)) // -2
	fmt.Println(math.Min(2, -3)) // -3

	fmt.Println(math.Max(1, 2))  // 2
	fmt.Println(math.Max(-2, 2)) // 2
	fmt.Println(math.Max(2, -3)) // 2
}
