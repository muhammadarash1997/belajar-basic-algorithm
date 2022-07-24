package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Repeat("go", 2)
	fmt.Println(s)

	s = strings.Repeat("lang", 1)
	fmt.Println(s)

	s = strings.Repeat("HI", 4)
	fmt.Println(s)

	s = strings.Repeat("BI", 3)
	fmt.Println(s)
}

// OUTPUT
// gogo
// lang
// HIHIHIHI
// BIBIBI