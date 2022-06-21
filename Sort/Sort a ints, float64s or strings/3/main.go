package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"dimas", "agus", "ramdan", "joko"}

	fmt.Println("Before", slice)
	sort.Strings(slice)
	fmt.Println("After", slice)
}