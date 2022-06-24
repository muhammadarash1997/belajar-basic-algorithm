package main

import (
    "fmt"
    "sort"
)

func main() {
    data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
	
	// since sort.Search using binary search, we need to sort data first
    sort.Ints(data)
    fmt.Println(data)

	// find x in data
    x := 8

	// sort.Search only allow us to use either >= or <=. If data is asceding order we use >= but if data is desceding order we use <=
    i := sort.Search(len(data), func(j int) bool { return data[j] >= x })
    if i >= len(data) || data[i] != x {
        // x is not present in data,
        // but i is the index where it would be inserted.
        fmt.Println("NO")
		return
    }
    fmt.Println("YES")
	return
}