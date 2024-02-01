package main

import "fmt"

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func twoReturns(n1, n2 int8) (int8, int8) {
	sumResult := sum(n1, n2)
	sub := n1 - n2

	return sumResult, sub
}

func main() {
	var txt = func(txt string) string {
		return "Text" + "" + txt
	}

	result := sum(1, 1) 
	fmt.Print(result, "\n")
	fmt.Println(twoReturns(1, 1))
	fmt.Println(txt(""))
	fmt.Println(txt("test"))

	// lint error when not use destruct in two variables
	sumResult, subResult := twoReturns(5, 5)
	fmt.Println(sumResult, subResult)

	// ignore return variable
	sumRes, _ := twoReturns(10,5)
	fmt.Println(sumRes)
}
