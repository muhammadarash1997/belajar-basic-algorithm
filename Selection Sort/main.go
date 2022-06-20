package main

import "fmt"

func main() {
	list := []int{2, 5, 3, 6, 8, 3, 1, 3, 5}
	fmt.Println(list)

	for i := range list {
		minIndex := i
		for j := i; j < len(list)-1; j++ {
			if list[j+1] < list[minIndex] {
				minIndex = j + 1
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}

	fmt.Println(list)
}

// Concept of Selection Sort is finding smallest element in array and moving it to front position but front position is always incremented by 1