package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool\n")
		case int:
			fmt.Printf("I'm an int\n")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}

	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello!")
}
