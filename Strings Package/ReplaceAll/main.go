package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.ReplaceAll("Go Golang G", "G", "EO")
	fmt.Println(s)

	s = strings.ReplaceAll("Go Golang G", "G", "EO")
	fmt.Println(s)

	s = strings.ReplaceAll("Go Golang G", "G", "EO")
	fmt.Println(s)

	s = strings.ReplaceAll("Go Golang G", "G", "EO")
	fmt.Println(s)

	s = strings.ReplaceAll("Go Golang G", "G", "")
	fmt.Println(s)

	s = strings.ReplaceAll("Go Golang G", "", "EO")
	fmt.Println(s)
}

// OUTPUT
// EOo EOolang EO
// EOo EOolang EO
// EOo EOolang EO
// EOo EOolang EO
// o olang 
// EOGEOoEO EOGEOoEOlEOaEOnEOgEO EOGEO