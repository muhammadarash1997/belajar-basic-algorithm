package main

import (
	"fmt"
	"sort"
	"strings"
)

type Nodes []Node

type Node struct {
	key string
	value int
}

func main() {
	// iNodes := Nodes{
	// 	{"Alice", 2},
	// 	{"Cecil", 1},
	// 	{"Mark", 2},
	// 	{"Biky", 2},
	// 	{"Bob", 3},
	// }

	myMap := map[string]int{"Alice": 2, "Cecil": 1, "Mark": 2, "Biky": 2, "Bob": 3}

	iNodes := Nodes{}

	for k, v := range myMap {
		iNodes = append(iNodes, Node{k, v})
	}

	fmt.Println(iNodes)

	sort.Slice(iNodes, func(i, j int) bool {
		if iNodes[i].value < iNodes[j].value {
			return true
		}
		if iNodes[i].value > iNodes[j].value {
			return false
		}
		return iNodes[i].key < iNodes[j].key
	})

	fmt.Println(iNodes)
}
