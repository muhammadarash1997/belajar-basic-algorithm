package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Title("Hello google")
	fmt.Println(s)

	s = strings.Title("Hello go golang")
	fmt.Println(s)

	s = strings.Title("Hey Jon !!!")
	fmt.Println(s)

	s = strings.Title("Teremenm fr # $kk ")
	fmt.Println(s)
}

// OUTPUT
// Hello Google
// Hello Go Golang
// Hey Jon !!!
// Teremenm Fr # $Kk
