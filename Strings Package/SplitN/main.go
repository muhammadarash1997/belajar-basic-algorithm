package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(s)

	s = strings.SplitN("Hello Go Golang", "o", 1)
	fmt.Println(s)

	s = strings.SplitN("abc", "", 3)
	fmt.Println(s)

	s = strings.SplitN("a,b,c", ",", 1)
	fmt.Println(s)

	s = strings.SplitN("xyz", ",", 2)
	fmt.Println(s)

	s = strings.SplitN("xyz", "p", 3)
	fmt.Println(s)
}
