package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	// 'i' does not present in 'hello'
	fmt.Println(strings.ContainsAny("hello", "i"))

	// 'i' in 'ui' present in 'fail'
	fmt.Println(strings.ContainsAny("fail", "ui"))

	// 'u' in 'ui' present in 'urge'
	fmt.Println(strings.ContainsAny("urge", "ui"))

	// 'u' and 'i' in 'ui' present in 'failure'
	fmt.Println(strings.ContainsAny("failure", "ui"))

	fmt.Println(strings.ContainsAny("foo", ""))

	fmt.Println(strings.ContainsAny("", ""))
}

// OUTPUT
// false
// true
// true
// true
// false
// false