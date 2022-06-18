package main

import (
	"fmt"
	"sort"
)

func main() {

	key := make(map[string]int)

	key["n"] = 1
	key["s"] = 2
	key["g"] = 1
	key["j"] = 2
	key["o"] = 3
	key["e"] = 1

	keys := []string{}

	for i := range key {
		keys = append(keys, i)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		// return keys[i] < keys[j] && key[keys[i]] < key[keys[j]]
		return key[keys[i]] < key[keys[j]]
	})
	fmt.Println(keys)

	for _, d := range keys {
		fmt.Println(d, key[d])
	}
}

// Program di atas bertujuan untuk mendisplay berdasarkan urutan key dan juga urutan value