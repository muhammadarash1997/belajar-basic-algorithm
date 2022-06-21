package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	fmt.Println(strings.Split("x,y,z", ","))

	fmt.Println(strings.Split("Hello Hii techieindoor", "i "))

	fmt.Println(strings.Split(" pqr", ""))

	fmt.Println(strings.Split("", "Hello O'Phiphi"))

	fmt.Println(strings.Split("hello", ""))

}

// OUTPUT
// [x y z]
// [Hello Hi techieindoor]
// [  p q r]
// []
// %q
// [h e l l o]