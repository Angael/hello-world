package pointers

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Run() {
	var p1 Person = Person{Name: "John", Age: 30}
	p2 := &p1

	p2.Name = "Doe"
	p2.Age = 40

	fmt.Println("Person 1: ", p1)
	fmt.Println("Person 2: ", *p2)

	p1.Name = "Alice"
	p1.Age = 25

	fmt.Println("Person 1: ", p1)
	fmt.Println("Person 2: ", *p2)
}
