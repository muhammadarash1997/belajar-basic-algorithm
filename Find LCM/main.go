package main

import (
	"fmt"
)

func main() {
	input := []int{25, 10}
	fmt.Println(FindLCM(input...))
}

func FindLCM(input ...int) int {
	product := 1
	for _, v := range input {
		product *= v
	}
	return product / FindGCD(input...)
}

func FindGCD(input ...int) int {
	GCD := 0
	keys := make(map[int]int)
	for _, d := range input {

		for i := 1; i <= d; i++ {

			if d%i == 0 {
				keys[i]++
				if keys[i] == len(input) {
					if keys[i] > GCD {
						GCD = i
					}
				}
			}
		}
	}
	return GCD
}