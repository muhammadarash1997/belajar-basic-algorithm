package main

import (
    "fmt"
    "strings"
)
func main() {
    s := strings.Join([]string{"hello", "world"}, ", ")
    fmt.Println(s)
}

// OUTPUT
// hello, world