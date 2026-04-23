package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func process(work <-chan string, fin chan<- string) {
	var b bytes.Buffer
	for {
		if msg, notClosed := <-work; notClosed {
			fmt.Printf("Processing: %s\n", msg)
		} else {
			fmt.Println("channel closed, finishing up")
			fin <- b.String()
			return
		}
	}
}

func writer(msgChan chan<- string, msg string) {
	msgChan <- msg
}

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func test1(c chan<- string) {
	for {
		pause()
		c <- "Hello from test1"
	}
}

func test2(c chan<- string) {
	for {
		pause()
		c <- "Hello from test2"
	}
}

func reader(msgChan <-chan string) {
	msg := <-msgChan
	fmt.Println("Received message:", msg)
}

func main() {
	msgChan := make(chan string)

	msgBuffer := make(chan string, 4) // Buffered channel with capacity of 4

	newMsgChan := make(chan string, 1)

	c1 := make(chan string)
	c2 := make(chan string)

	work := make(chan string, 3)
	fin := make(chan string)

	cs := make(chan string, 3)

	go func() {
		cs <- "Hello from goroutine 1"
		cs <- "Hello from goroutine 2"
		cs <- "Hello from goroutine 3"
		close(cs)
	}()

	for msg := range cs {
		fmt.Println("Received from cs:", msg)
	}

	go process(work, fin)
	word := "hello world"

	for i := 0; i < len(word); i++ {
		letter := string(word[i])
		work <- letter
		fmt.Printf("Sent: %s\n", letter)
	}
	close(work)

	fmt.Printf("Final result: %s\n", <-fin)
	test1(c1)
	test2(c2)
	timeout := time.After(3 * time.Second)
	for {
		select {
		case msg := <-c1:
			fmt.Println("Received from test1:", msg)
		case msg := <-c2:
			fmt.Println("Received from test2:", msg)
		case <-timeout:
			fmt.Println("Timeout waiting for messages")
			return
		}
	}

	go func(channel chan string) {
		time.Sleep(1 * time.Second)
		channel <- "Hello from the goroutine"
	}(c1)

	select {

	case msg2 := <-c1:
		fmt.Println("Received from goroutine:", msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout waiting for message")
	}
	go reader(newMsgChan)

	for i := 1; i <= 4; i++ {
		writer(newMsgChan, fmt.Sprintf("Hello from goroutine %d", i))
	}

	time.Sleep(1 * time.Second) // Give some time for the reader to process messages

	go func() {

		_ = <-msgBuffer             // This will not block because the channel is buffered
		time.Sleep(2 * time.Second) // Simulate some processing time
	}()

	writer := func() {
		for i := 1; i <= 4; i++ {
			msgBuffer <- fmt.Sprintf("Hello from goroutine %d", i) // This will not block until the buffer is full
		}
	}

	writer()

	go func() {
		msgChan <- "Hello from goroutine 1"
	}()

	go func() {
		msgChan <- "Hello from goroutine 2"
	}()

	go func() {
		msgChan <- "Hello from goroutine 3"
	}()

	go func() {
		msgChan <- "Hello from goroutine 4"
	}()

	for i := 0; i < 4; i++ {
		fmt.Println(<-msgChan)
	}

	close(msgChan)
}
