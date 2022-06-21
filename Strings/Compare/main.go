package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	fmt.Println(strings.Compare("Hello", "Hi"))

	fmt.Println(strings.Compare("Hello", "Hello"))

	fmt.Println(strings.Compare("Hi", "Hello"))

}

// OUTPUT
// -1
// 0
// 1