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

func pathfinderDamage(attackBonus int, damageDice int, damageBonus int, critRange int, rolls int, ac int) float64 {
	var totalDamage int
	for i := 0; i < rolls; i++ {
		roll := RollDice(20)
		if roll >= critRange {
			// Critical hit
			totalDamage += RollDice(damageDice) + damageBonus + RollDice(damageDice) + damageBonus
		} else if roll+attackBonus >= ac { // Assuming AC is 10 for simplicity
			// Normal hit
			totalDamage += RollDice(damageDice) + damageBonus
		}
		// else miss, do nothing
	}

	// Round to 2 decimal places
	rounded := float64(totalDamage) / float64(rolls)
	return float64(int(rounded*1000)) / 1000
}

func Run() {
	fmt.Println("Rolling a 6-sided dice:", RollDice(6))
	fmt.Println("Rolling a 20-sided dice:", RollDice(20))
	fmt.Println("Average roll of a 6-sided dice (100 rolls):", TestAverageRoll(6, 100))

	rolls := 1000000
	fmt.Println("--- Pathfinder Damage Comparison (", rolls, " rolls) ---")

	// test for AC from 7 to 20
	for ac := 7; ac <= 20; ac++ {
		fmt.Println("AC: ", ac)

		// (attack +6, damage 2d6 + 4, critical d19-20)
		greatSword := pathfinderDamage(6, 6, 4, 19, rolls, ac)

		// (attack +4, damage 1d8 + 3, critical d19-20)
		longSword := pathfinderDamage(4, 8, 3, 19, rolls, ac)

		// (attack +3, damage 1d6 + 1, critical d20)
		caseLightMace := pathfinderDamage(3, 6, 1, 20, rolls, ac)

		// (attack +4, damage 1d4 + 1, critical d19-20)
		caseDagger := pathfinderDamage(4, 4, 1, 19, rolls, ac)

		fmt.Println("  Longsword: ", longSword)
		fmt.Println("  Light Mace: ", caseLightMace)
		fmt.Println("  Dagger: ", caseDagger)

		fmt.Println("  Greatsword: ", greatSword)

		roundedLongSwordAndLightMace := float64(int((longSword+caseLightMace)*1000)) / 1000
		fmt.Println("  Longsword + Light Mace: ", roundedLongSwordAndLightMace)

		roundedLongSwordAndDagger := float64(int((longSword+caseDagger)*1000)) / 1000
		fmt.Println("  Longsword + Dagger: ", roundedLongSwordAndDagger)

	}
}
