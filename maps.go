package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	clear(m)
	fmt.Println("The value:", m["Answer"])

	_, ok = m["Answer"]
	fmt.Println("Present?", ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map:", n2)
	if maps.Equal(n, n2) {
		fmt.Println("maps are equal")
	}
}
