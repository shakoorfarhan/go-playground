package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function.
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	for n := range 6 {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
	fmt.Println(sum)
}
