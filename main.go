package main

import (
	"flag"
	"fmt"

	functions "example/hello/lessons/00_functions"
	loops "example/hello/lessons/01_loops"
	colors "example/hello/lessons/02_colors"
	channels "example/hello/lessons/03_channels"
	pointers "example/hello/lessons/04_pointers"
	goroutines "example/hello/lessons/05_goroutines"
	randomdice "example/hello/lessons/06_randomdice"
	manythreads "example/hello/lessons/07_manythreads"
	arrays "example/hello/lessons/08_arrays"
)

func main() {
	lesson := flag.Int("lesson", -1, "Run a specific lesson (0-6), -1 for all")
	flag.Parse()

	if *lesson == -1 || *lesson == 0 {
		fmt.Println("--- Functions Lesson ---")
		functions.Run()
	}

	if *lesson == -1 || *lesson == 1 {
		fmt.Println("--- Loops Lesson ---")
		loops.Run()
	}

	if *lesson == -1 || *lesson == 2 {
		fmt.Println("--- Colors Lesson ---")
		colors.Run()
	}

	if *lesson == -1 || *lesson == 3 {
		fmt.Println("--- Channels Lesson ---")
		channels.Run()
	}

	if *lesson == -1 || *lesson == 4 {
		fmt.Println("--- Pointers ---")
		pointers.Run()
	}

	if *lesson == -1 || *lesson == 5 {
		fmt.Println("--- Goroutines ---")
		goroutines.Run()
	}

	if *lesson == -1 || *lesson == 6 {
		fmt.Println("--- Random Dice ---")
		randomdice.Run()
	}

	if *lesson == -1 || *lesson == 7 {
		fmt.Println("--- Many threads sync ---")
		manythreads.Run()
	}

	if *lesson == -1 || *lesson == 8 {
		fmt.Println("--- Arrays ---")
		arrays.Run()
	}

}
