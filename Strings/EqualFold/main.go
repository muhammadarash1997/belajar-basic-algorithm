package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	fmt.Println(strings.EqualFold("Go", "go"))

	fmt.Println(strings.EqualFold("Go", "lang"))

}

// OUTPUT
// true
// false