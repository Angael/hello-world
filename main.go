package main

import (
	"example/hello/lessons/00_functions"
	"example/hello/lessons/01_loops"
	"example/hello/lessons/02_colors"
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
}
