package main

import "fmt"

func SliceIndex[S ~[]E, E comparable](s S, e E) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func main() {
	s := []string{"apple", "banana", "cherry"}
	fmt.Println(SliceIndex(s, "banana")) // Output: 1
	fmt.Println(SliceIndex(s, "grape"))  // Output: -1

	n := []int{10, 20, 30}
	fmt.Println(SliceIndex(n, 20)) // Output: 1
	fmt.Println(SliceIndex(n, 40)) // Output: -1
}
