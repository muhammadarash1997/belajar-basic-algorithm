package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	K := 5 // variabel want to be added

	afterAppend := AppendToFirst(K, A)
	fmt.Println(afterAppend)
}

func AppendToFirst(f int, s []int) []int {
	l := []int{f}
	l = append(l, s...)
	return l
}