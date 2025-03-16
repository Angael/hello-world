package randomdice

import (
	"fmt"
	"math/rand"
)

func RollDice(n int) int {
	return rand.Intn(n) + 1
}

func TestAverageRoll(n int, rolls int) float64 {
	var sum int
	for i := 0; i < rolls; i++ {
		sum += RollDice(n)
	}
	return float64(sum) / float64(rolls)
}

type Weapon struct {
	Name        string
	AttackBonus int
	DamageDice  int
	DamageBonus int
	CritRange   int
}

func (w Weapon) AverageDamage(rolls int, ac int) float64 {
	var totalDamage int
	for i := 0; i < rolls; i++ {
		roll := RollDice(20)
		if roll >= w.CritRange {
			// Critical hit
			totalDamage += RollDice(w.DamageDice) + w.DamageBonus + RollDice(w.DamageDice) + w.DamageBonus
		} else if roll+w.AttackBonus >= ac {
			// Normal hit
			totalDamage += RollDice(w.DamageDice) + w.DamageBonus
		}
		// else miss, do nothing
	}

	// Round to 3 decimal places
	rounded := float64(totalDamage) / float64(rolls)
	return float64(int(rounded*1000)) / 1000
}

func Run() {
	fmt.Println("Rolling a 6-sided dice:", RollDice(6))
	fmt.Println("Rolling a 20-sided dice:", RollDice(20))
	fmt.Println("Average roll of a 6-sided dice (100 rolls):", TestAverageRoll(6, 100))

	rolls := 1000000
	fmt.Println("--- Pathfinder Damage Comparison (", rolls, " rolls) ---")

	// Define weapons
	greatsword := Weapon{Name: "Greatsword", AttackBonus: 6, DamageDice: 6, DamageBonus: 4, CritRange: 19}
	longsword := Weapon{Name: "Longsword", AttackBonus: 4, DamageDice: 8, DamageBonus: 3, CritRange: 19}
	lightMace := Weapon{Name: "Light Mace", AttackBonus: 3, DamageDice: 6, DamageBonus: 1, CritRange: 20}
	dagger := Weapon{Name: "Dagger", AttackBonus: 4, DamageDice: 4, DamageBonus: 1, CritRange: 19}

	// test for AC 5, 10, 15
	for _, ac := range []int{5, 10, 15} {
		fmt.Println("AC: ", ac)

		// Calculate average damage
		greatswordDamage := greatsword.AverageDamage(rolls, ac)
		longswordDamage := longsword.AverageDamage(rolls, ac)
		lightMaceDamage := lightMace.AverageDamage(rolls, ac)
		daggerDamage := dagger.AverageDamage(rolls, ac)

		fmt.Println("  Greatsword: ", greatswordDamage)

		roundedLongSwordAndLightMace := float64(int((longswordDamage+lightMaceDamage)*1000)) / 1000
		fmt.Println("  Longsword + Light Mace: ", roundedLongSwordAndLightMace)

		roundedLongSwordAndDagger := float64(int((longswordDamage+daggerDamage)*1000)) / 1000
		fmt.Println("  Longsword + Dagger: ", roundedLongSwordAndDagger)
	}
}
