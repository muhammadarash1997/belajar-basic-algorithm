package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	a := "Was it a car or a cat I saw?"
	fmt.Println(Solution(a))
}

func Solution(str string) bool {
	// Your code starts here.

	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	r, err := regexp.Compile(`[^\w]`)
	if err != nil {
		panic(err)
	}
	str = r.ReplaceAllString(str, "")

	return isPalindrome(str)
}

func isPalindrome(str string) bool {
	var tmp []byte
	for i := len(str) - 1; i >= 0; i-- {
		tmp = append(tmp, str[i])
	}

	if string(tmp) == str {
		return true
	}

	return false
}
