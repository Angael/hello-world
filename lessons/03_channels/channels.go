package channels

import "fmt"

func Run() {
	// Create a channel
	ch := make(chan string)

	// Send a value to the channel in a goroutine
	go func() {
		ch <- "Hello, Channel!"
	}()

	// Receive the value from the channel
	msg := <-ch

	fmt.Println("Received message:", msg)

	// Closing channels
	close(ch)

	// Example of sending and receiving multiple values
	intChan := make(chan int, 3) // Buffered channel with capacity 3
	intChan <- 1
	intChan <- 2
	intChan <- 3
	close(intChan) // Important to close the channel after sending all values

	// Receiving values from a closed channel
	for val := range intChan {
		fmt.Println("Received int:", val)
	}
}
