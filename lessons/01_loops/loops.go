package loops

import "fmt"

func Run() {
	numOfDays := 365
	numOfDaysInWeek := 7
	numOfDaysInMonth := 30

	fmt.Println("Number of days in a year: ", numOfDays)
	fmt.Println("Number of days in a week: ", numOfDaysInWeek)
	fmt.Println("Number of days in a month: ", numOfDaysInMonth)

	var numOfFridays int
	for i := 0; i < numOfDays; i++ {
		if i%numOfDaysInWeek == 5 {
			numOfFridays++
			fmt.Println("Friday nr: ", numOfFridays, "Day: ", i+1)
		}

	}

	fmt.Println("Number of Fridays in a year: ", numOfFridays)
}
