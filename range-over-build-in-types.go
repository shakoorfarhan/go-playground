package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
			break
		}
	}

	k := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range k {
		fmt.Printf("key:%s value:%d\n", key, value)
	}

	for d := range k {
		fmt.Println("key:", d)
	}

	for i, c := range "go" {
		fmt.Printf("index:%d char:%c\n", i, c)
	}
}
