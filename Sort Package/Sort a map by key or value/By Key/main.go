package main

import (
	"fmt"
	"sort"
)

// Sort by key
func main() {
	myMap := map[string]int{"Alice": 2, "Cecil": 1, "Mark": 2, "Bob": 3}

	keys := []string{}
	for i := range myMap {
		keys = append(keys, i)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, myMap[k])
	}
}

// OUTPUT
// Alice 2
// Bob 3
// Cecil 1
// Mark 2