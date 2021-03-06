package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Replace("Go Golang G", "G", "EO", 2)
	fmt.Println(s)

	s = strings.Replace("Go Golang G", "G", "EO", -1)
	fmt.Println(s)

	s = strings.Replace("Go Golang G", "G", "EO", 1)
	fmt.Println(s)

	s = strings.Replace("Go Golang G", "G", "EO", 0)
	fmt.Println(s)

	s = strings.Replace("Go Golang G", "G", "", 2)
	fmt.Println(s)

	s = strings.Replace("Go Golang G", "", "EO", 1)
	fmt.Println(s)
}

// OUTPUT
// EOo EOolang G
// EOo EOolang EO
// EOo Golang G
// Go Golang G
// o olang G
// EOGo Golang G