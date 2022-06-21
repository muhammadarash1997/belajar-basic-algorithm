package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"dimas", "agus", "ramdan", "joko"}

	fmt.Println("Before:", slice)
	sort.Strings(slice)
	fmt.Println("After:", slice)
}

// OUTPUT
// Before: [dimas agus ramdan joko]
// After: [agus dimas joko ramdan]