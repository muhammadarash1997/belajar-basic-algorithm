package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Greg", 17},
		{"Jenny", 26},
	}
	fmt.Println("Original:", people)

	// THREE WAYS TO SORT A SLICE

	// FIRST
	// Sort sorts data in ascending order as determined by the Less method. It makes one call to data.Len
	// to determine n and O(n*log(n)) calls to data.Less and data.Swap. The sort is not guaranteed to be stable. 
	sort.Sort(ByAge(people))
	fmt.Println("Using Sort:", people)

	// SECOND
	// Slice sorts the slice x given the provided less function. It panics if x is not a slice.
	// The sort is not guaranteed to be stable: equal elements may be reversed from their original order. For a stable sort, use SliceStable. 
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("Using Slice:", people)

	// THIRD
	// SliceStable sorts the slice x using the provided less function, keeping equal elements in their original order. It panics if x is not a slice.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("Using Slice:", people)
}
