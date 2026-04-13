package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{"bob", 20})
	fmt.Println(Person{name: "alice", age: 30})
	fmt.Println(Person{name: "fred"})

	fmt.Println(&Person{name: "ann", age: 40})
	fmt.Println(newPerson("jon"))

	s := Person{name: "sds", age: 33}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}

	fmt.Println(dog)
}
