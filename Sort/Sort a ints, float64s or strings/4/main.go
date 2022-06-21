package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int32{4, 2, 3, 2, 1}

	fmt.Println("Before", slice)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Println("After", slice)
}
