package main

import "fmt"

func func1() {
	fmt.Println("Função 1")
}
func func2() {
	fmt.Println("Função 2")
}
func func3() {
	fmt.Println("Função 3")
}
func funcDefer1() {
	fmt.Println("Função defered 1")
}
func funcDefer2() {
	fmt.Println("Função defered 2")
}

func calcApprovement(x, y float32) bool {
	defer fmt.Println("Your approvement was calculated")

	fmt.Println("Your approvement will be calculated")
	mean := (x + y) / 2

	return mean >= 7
}

func main() {
	func1()
	func2()
	func3()

	fmt.Println("----DEFER === adiar ----")
	defer func1()
	defer func2()
	func3()

	fmt.Println("--------")
	defer funcDefer2()
	defer funcDefer1()
	func3()

	// DEFER IS LIFO IN CALLSTACK - LAST IN FIRST OUT

	fmt.Println("---Defer inside function---")
	fmt.Println(calcApprovement(5, 7))

	fmt.Println("---Defer ON MAIN---")
}
