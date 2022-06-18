package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Before swap: %v\n", slice)

	swap := reflect.Swapper(slice)

	swap(2, 4)

	fmt.Printf("After swap: %v\n", slice)
}
