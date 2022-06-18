package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{-23, 567, -34, 67, 0, 12, -5}

	fmt.Println("Before", slice)

	sort.Ints(slice)

	fmt.Println("After", slice)
}
