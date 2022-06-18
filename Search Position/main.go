package main

import "fmt"

func SearchPosition(s []int, v int) int {
	for i, d := range s {
		if d == v {
			return i + 1
		}
	}
	return -1
}

func main() {
	slice := []int{5, 4, 3, 2, 1}

	position := SearchPosition(slice, 5)

	fmt.Println(position)
}
