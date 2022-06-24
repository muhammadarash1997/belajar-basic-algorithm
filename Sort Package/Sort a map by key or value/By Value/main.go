package main

import (
	"fmt"
	"sort"
)

type Node struct {
	key string
	value int
}

// Sort by key and value
// Sort by value first then by key
func main() {

	// Mapping map to struct
	myMap := map[string]int{"Biky": 2, "Cecil": 1, "Mark": 2, "Alice": 2, "Bob": 3}
	iNodes := []Node{}
	for k, v := range myMap {
		iNodes = append(iNodes, Node{k, v})
	}
	fmt.Println("Before:", iNodes)

	sort.Slice(iNodes, func(i, j int) bool {
		return iNodes[i].value < iNodes[j].value
	})
	fmt.Println("After", iNodes)
}

// OUTPUT
// Before: [{Biky 2} {Bob 3} {Alice 2} {Cecil 1} {Mark 2}]
// After [{Cecil 1} {Alice 2} {Biky 2} {Mark 2} {Bob 3}]