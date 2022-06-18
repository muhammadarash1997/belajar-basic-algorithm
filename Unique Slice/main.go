package main

import "fmt"

func Unique(s []int) []int {
	m := make(map[int]bool)
	l := []int{}
	for _, d := range s {
		v := m[d]
		if !v {
			m[d] = true
			l = append(l, d)
		}
	}
	return l
}

func main() {
	slice := []int{1, 2, 3, 3, 5, 5, 5, 6}

	uniqueSlice := Unique(slice)
	
	fmt.Println(uniqueSlice)
}