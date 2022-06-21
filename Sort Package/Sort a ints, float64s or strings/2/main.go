package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []float64{1.2, 5.3, 8.5, 3.4, 2.5, 5.3}

	fmt.Println("Before:", slice)
	sort.Float64s(slice)
	fmt.Println("After:", slice)
}

// OUTPUT
// Before: [1.2 5.3 8.5 3.4 2.5 5.3]
// After: [1.2 2.5 3.4 5.3 5.3 8.5]