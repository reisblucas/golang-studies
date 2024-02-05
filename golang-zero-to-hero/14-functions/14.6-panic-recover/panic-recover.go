package main

import "fmt"

func showDeferAfterPanic() {
	// fmt.Println("--------\nI'M DEFERED BEFORE PANIC\n--------")
	if r := recover(); r != nil {
		fmt.Println("Your execution was recovered with success!")
	}
}

func calcApprovement(x, y float32) bool {
	defer showDeferAfterPanic()
	defer fmt.Println("Your approvement was calculated")

	fmt.Println("Your approvement will be calculated")
	mean := (x + y) / 2

	if mean == 7 {
		panic("The mean is equal to 7")
	}

	return mean > 7
}

func main() {
	fmt.Println(calcApprovement(6, 7))
	fmt.Println("Pre panic")
	fmt.Println(calcApprovement(7, 7))
	fmt.Println("Pos panic")
}
