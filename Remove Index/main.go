package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60}
	index := 5	// index want to remove
	sliceAfterRemove := RemoveIndex(slice, index)
	fmt.Println(sliceAfterRemove)
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}