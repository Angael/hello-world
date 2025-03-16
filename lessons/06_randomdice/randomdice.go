package randomdice

import (
	"fmt"
	"math/rand"
	"time"
)

func RollDice(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n) + 1
}

func Run() {
	fmt.Println("Rolling a 6-sided dice:", RollDice(6))
	fmt.Println("Rolling a 20-sided dice:", RollDice(20))
}
