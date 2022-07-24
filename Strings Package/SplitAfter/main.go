package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.SplitAfter("a,b,c", ",")
	fmt.Println(s)

	s = strings.SplitAfter("Hello Go Golang", "o")
	fmt.Println(s)

	s = strings.SplitAfter("abc", "")
	fmt.Println(s)

	s = strings.SplitAfter("a,b,c", ",")
	fmt.Println(s)

	s = strings.SplitAfter("xyz", ",")
	fmt.Println(s)

	s = strings.SplitAfter("xyz", "p")
	fmt.Println(s)
}

// OUTPUT
// [a, b, c]
// [Hello  Go  Go lang]
// [a b c]
// [a, b, c]
// [xyz]
// [xyz]