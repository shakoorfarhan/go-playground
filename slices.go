package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	s := make([]int, 5)

	s[0] = 1
	s[1] = 2

	s[2] = 1
	s[3] = 2
	s[4] = 1

	s = append(s, 0)
	fmt.Println("set:", s)
	fmt.Println("get:", s[4])
	fmt.Println("len:", len(s))
	c := make([]int, len(s))
	copy(c, s)

	fmt.Println("c", c)
	fmt.Println("s:", s)

	fmt.Println("2", s[:2])
	fmt.Println("3", s[2:])
	fmt.Println("4", s[1:4])

	twoD := make([][]int, 3)

	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD)
}
