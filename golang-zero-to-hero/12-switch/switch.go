package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	default:
		return "Invalid number"
	}
}

func weekDayDifferentSwitch(number int) string {
	switch {
	case number == 0:
		return "Sunday"
	case number == 1:
		return "Monday"
	default:
		return "Invalid number"
	}
}

func weekDayWithFallthrough(number int) string {
	var weekDay string

	switch {
	case number == 0:
		weekDay = "Sunday"
		fallthrough // skip to the next loop without case validation
	case number == 1:
		weekDay = "Monday"
	default:
		weekDay = "Invalid number"
	}

	return weekDay
}

func main() {
	fmt.Println(weekDay(1))
	fmt.Println(weekDay(3))
	fmt.Println(weekDayWithFallthrough(0))
	fmt.Println(weekDayWithFallthrough(1))
}
