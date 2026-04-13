package main

import "fmt"

type base struct {
	name string
}

func (b base) getName() string {
	return b.name
}

type derived struct {
	base
	age int
}

func main() {
	d := derived{
		base: base{name: "John"},
		age:  30,
	}
	fmt.Println(d.getName()) // Output: John
	fmt.Println(d.age)       // Output: 30
}
