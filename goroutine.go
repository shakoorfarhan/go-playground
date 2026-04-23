package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	pause()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4) // We will launch 4 goroutines

	go func(msg string) {
		defer wg.Done()
		pause()
		fmt.Println(msg)

	}("Hello from goroutine 4")

	go sendMessage("Hello from goroutine 1", &wg)
	go sendMessage("Hello from goroutine 2", &wg)
	go sendMessage("Hello from goroutine 3", &wg)
	wg.Wait() // Wait for all goroutines to finish
}
