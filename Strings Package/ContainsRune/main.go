package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
	// Finds whether a string contains a particular Unicode code
	// point.

	// The code point for the lowercase letter "a", for example,
	// is 97.

	fmt.Println(strings.ContainsRune("aadlik", 97))

	fmt.Println(strings.ContainsRune("peek", 97))

}

// OUTPUT
// true
// false