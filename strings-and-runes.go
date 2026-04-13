package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Strings are a slice of bytes
	s := "Hello, 世界"
	fmt.Println("String:", s)
	fmt.Println("Length in bytes:", len(s))

	// To get the number of runes (Unicode code points)
	runeCount := utf8.RuneCountInString(s)
	fmt.Println("Length in runes:", runeCount)

	// Iterating over a string with range gives you runes
	fmt.Println("Runes in the string:")
	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}
}
