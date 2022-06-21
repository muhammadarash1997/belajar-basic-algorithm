package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
 
  s := strings.ToLower("Go GoLang")
  fmt.Println(s)
 
  s = strings.ToLower("go golang")
  fmt.Println(s)
 
  s = strings.ToLower("Go GoLang 123 #")
  fmt.Println(s)
 
  s = strings.ToLower("0 9 # Go GoLang ")
  fmt.Println(s)
}

// OUTPUT
// go golang
// go golang
// go golang 123 #
// 0 9 # go golang 