package arrays

import (
	"fmt"
	"math/rand"
)

type Person struct {
	Index   int
	Name    string
	Attack  int
	Defense int
	Health  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s [%d] (%d/%d)", p.Name, p.Health, p.Attack, p.Defense)
}

func (p *Person) AttackEnemy(enemy *Person) {
	// minimal damage is 0
	if p.Attack < enemy.Defense {
		return
	}

	// 50 defense is 50% damage reduction
	enemy.Health -= p.Attack * (100 - enemy.Defense) / 100

	if enemy.Health <= 0 {
		enemy.Health = 0
		fmt.Println("Person", enemy, "is dead")
	}
}

func cleanupDead(people []Person) []Person {
	var survivors []Person
	var deadPeople int = 0
	for _, person := range people {
		if person.Health > 0 {
			survivors = append(survivors, person)
		} else {
			deadPeople++
		}
	}

	fmt.Println("Dead people:", deadPeople)

	return survivors
}

func Run() {
	names := []string{"Albert", "Bob", "Charlie", "David", "Eve", "Frank", "George", "Helen", "Ivan", "Jack", "Kathy", "Lenny", "Mandy", "Nancy", "Oscar", "Pam", "Quincy", "Randy", "Steve", "Tina", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zack"}
	var people []Person

	// Add some people
	for i := 0; i < 100000; i++ {
		people = append(people, Person{
			Index:   i,
			Name:    names[rand.Intn(len(names))],
			Attack:  rand.Intn(200),
			Defense: rand.Intn(80),
			Health:  100,
		})
	}

	rounds := 100
	for j := 0; j < rounds; j++ {
		fmt.Println("  [ Round", j, "]")

		for i, person := range people {
			foeIndex := rand.Intn(len(people))

			if i == foeIndex {
				fmt.Println("Person", person, "was confused")

			} else {
				foe := &people[foeIndex]
				fmt.Println(i, person, "attacks", foe)
				person.AttackEnemy(foe)
			}
		}

		people = cleanupDead(people)
		if len(people) <= 1 {
			break
		}
	}

	fmt.Println("Survivors:")
	for _, person := range people {
		fmt.Println(person)
	}
}
