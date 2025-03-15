package goroutines

import (
	"fmt"
	"time"
)

func AsyncStuff(i *int) {
	fmt.Println("Async stuff")
	time.Sleep(500 * time.Millisecond)
	*i = 1
	fmt.Println("Async stuff done")
}

func Run() {
	i := 0

	fmt.Println("Async stuff, i = ", i)
	go AsyncStuff(&i)
	fmt.Println("Async stuff started, i = ", i)

	// sleep
	time.Sleep(1 * time.Second)

	fmt.Println("Async stuff finished, i = ", i)
}
