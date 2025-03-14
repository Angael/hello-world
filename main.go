package main

import (
	functions "example/hello/lessons/00_functions"
	loops "example/hello/lessons/01_loops"
	colors "example/hello/lessons/02_colors"
	pointers "example/hello/lessons/04_pointers"
	"fmt"

	"rsc.io/quote"
)

func main() {
	// Run lessons
	fmt.Println("--- Functions Lesson ---")
	functions.Run()

	fmt.Println("--- Loops Lesson ---")
	loops.Run()

	fmt.Println("--- Colors Lesson ---")
	colors.Run()

	fmt.Println("--- Quote ---")
	fmt.Println(quote.Go())

	fmt.Println("--- Pointers ---")
	pointers.Run()
}
