package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	var index int

	index = strings.Index("Gabrie", "br")
	fmt.Println(index)

	index = strings.Index("Gabrie", "bob")
	fmt.Println(index)
}

// OUTPUT
// 2
// -1