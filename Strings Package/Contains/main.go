package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	// 'foo' presents in 'animalsfood'
	fmt.Println(strings.Contains("animalsfood", "foo"))

	// 'bar' does not present in 'birdsfood'
	fmt.Println(strings.Contains("birdsfood", "bar"))

	fmt.Println(strings.Contains("animalsfood", ""))

	fmt.Println(strings.Contains("", ""))
}

// OUTPUT
// true
// false
// true
// true