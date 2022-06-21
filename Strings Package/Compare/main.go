package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	fmt.Println(strings.Compare("Hello", "Hi"))

	fmt.Println(strings.Compare("Hello", "Hello"))

	fmt.Println(strings.Compare("Hi", "Hello"))

	fmt.Println(strings.Compare("Hi", "hi"))

}

// OUTPUT
// -1
// 0
// 1
// -1