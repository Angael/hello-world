package functions

import "fmt"

func Add(x int8, y int8) (int8, error) {
	if x > 0 && y > 0 && x+y < 0 {
		return 0, fmt.Errorf("Overflow")
	}

	return x + y, nil
}

func Run() {
	// Try to overflow add int8
	var a int8 = 127
	var b int8 = 127
	result, err := Add(a, b)
	fmt.Println("Overflow: ", result, err)
}
