package main

import "fmt"

func main() {
	list := []int{2, 5, 3, 6, 8, 3, 1, 3, 5}
	fmt.Println(list)

	var loop bool = true
	for loop {
		loop = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				loop = true
			}
		}
	}

	fmt.Println(list)
}

// Bubble Sort can sort a list even if there is duplicate value