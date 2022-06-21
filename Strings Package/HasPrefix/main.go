package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	var res bool

	res = strings.HasPrefix("Techieindoor", "Techie")
	fmt.Println(res)

	res = strings.HasPrefix("Hello Techie!", "Hello")
	fmt.Println(res)

	res = strings.HasPrefix("Hello Geeks", "Geeks")
	fmt.Println(res)

	res = strings.HasPrefix("Hello Geeks", "")
	fmt.Println(res)
}