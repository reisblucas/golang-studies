package main

import "fmt"

func main() {
	var variable string = "Not infered variable"
	variable2 := "Infered variable"
	fmt.Println(variable)
	fmt.Println(variable2)

	var (
		variable3 string = "infered"
		variable4 string = "infered2"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variable 5", "Variable 6"
	fmt.Println(variable5, variable6)
	
	const constant1 string = "Constant 1"
	fmt.Println(constant1)
	
	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}