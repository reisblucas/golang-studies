package main

import "fmt"

func closure() func() {
	x := 1
	y := 2
	z := x + y

	function := func() {
		fmt.Printf("x + y equals to %d\n", z)
	}

	return function
}

func main() {
	x := 10
	fmt.Printf("CLOSURE: x equals to: %d\n", x)
	newfunc := closure()
	newfunc()
}
