package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func binary_search(array []int, value, low, high int) int {
	var mid int
	// if high == low && array[(low+high)/2] != value {
	if high == low && array[low] != value {
		return -1
	} else {
		mid = (low + high) / 2
		if array[mid] > value {
			return binary_search(array, value, low, mid)
		} else if array[mid] < value {
			return binary_search(array, value, mid+1, high)
		} else {
			return mid
		}
	}
}

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var value, answer int
	array := []int{}
	for i := 0; i < 10000; i++ {
		scanner.Scan()
		T, _ := strconv.Atoi(scanner.Text())
		array = append(array, T)
	}

	for i := 0; i < 10000; i++ {
		scanner.Scan()
		value, _ = strconv.Atoi(scanner.Text())
		answer = binary_search(array, value, 0, 9999)

		if _, err := os.Stat("output.in"); err == nil {
			f, err := os.OpenFile("output.in", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			s := fmt.Sprintf("%d\n", answer)
			_, err = f.WriteString(s)
			if err != nil {
				log.Fatal(err)
			}
		} else if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create("output.in")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			s := fmt.Sprintf("%d\n", answer)
			_, err = f.WriteString(s)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
