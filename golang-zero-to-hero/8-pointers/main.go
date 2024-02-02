package main

import "fmt"

func main() {
	var variable string = "Teste"
	var variable2 string = variable
	fmt.Println(variable, variable2)

	variable = "Uhu"
	fmt.Println(variable, variable2)

	// pointers is a memory refrerence
	var somevariable int
	var reference *int

	somevariable = 100
	reference = &somevariable
	// dereferencing or indirecting
	fmt.Println(somevariable, reference, *reference)

	somevariable = 150
	fmt.Println(somevariable, reference, *reference)
}
