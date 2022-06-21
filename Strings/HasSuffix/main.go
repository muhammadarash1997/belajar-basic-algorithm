package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	fmt.Println(strings.HasSuffix("TechieIndoor", "Indoor"))

	fmt.Println(strings.HasSuffix("TechieIndoor", "r"))

	fmt.Println(strings.HasSuffix("TechieIndoor", "Techie"))

	fmt.Println(strings.HasSuffix("TechieIndoor", ""))
}

// OUTPUT
// true
// true
// false
// true