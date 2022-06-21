package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{6, 2, 3, 2, 1}

	fmt.Println("Before", slice)
	sort.Ints(slice)
	fmt.Println("After", slice)
}
