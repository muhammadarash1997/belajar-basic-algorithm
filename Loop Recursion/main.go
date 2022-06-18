package main

import "fmt"

func main() {
	Por(2, 4)
}

func Por(max int, numberLoop int) int {
	if numberLoop == 0 {
		return 0
	}

	for i := 1; i <= max; i++ {
		fmt.Print(i)
	}
	Por(max, numberLoop-1)

	return 0
}

// Tujuan loop recursion adalah agar kita dapat menetukan ada berapa banyak jumlah entitas loop dalam sebuah loop, contoh :

// Contoh ingin ada 3 entitas loop dalam suatu loop
// For loop {

	// For loop {

		// For loop {

			// For loop {

			// }
		// }
	// }


// }

// Contoh ingin ada 2 entitas loop dalam suatu loop
// For loop {

	// For loop {

		// For loop {
	
		// }
	// }

// }