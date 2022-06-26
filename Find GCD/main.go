package main

import "fmt"

func main() {

	fmt.Println(FindGCD(10, 15))
}

func FindGCD(input ...int) int {

	list := []int{}
	keys := make(map[int]int)
	for _, d := range input {

		for i := 1; i <= d; i++ {

			if d%i == 0 {
				keys[i]++
				if keys[i] == len(input) {
					list = append(list, i)
				}
			}
		}
	}
	return list[len(list)-1]
}

// Contoh FPB dari 12, 16, 8 adalah 4
// Contoh FPB dari 10, 15 adalah 5