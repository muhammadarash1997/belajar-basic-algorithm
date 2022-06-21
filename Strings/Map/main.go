package main

import (
	"fmt"
	"strings"
)

func validate_fun(r rune) rune {

	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+10)%26

	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+10)%26
	}
	return r
}

func main() {

	str := "Hello all"
	s := strings.Map(validate_fun, str)
	fmt.Println(s)

	str = ""
	s = strings.Map(validate_fun, str)
	fmt.Println(s)

	str = "12345"
	s = strings.Map(validate_fun, str)
	fmt.Println(s)

	str = "-123.455"
	s = strings.Map(validate_fun, str)
	fmt.Println(s)

	str = "Hello all @@##$$"
	s = strings.Map(validate_fun, str)
	fmt.Println(s)
}
