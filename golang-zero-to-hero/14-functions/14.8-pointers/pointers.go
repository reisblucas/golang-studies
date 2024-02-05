package main

import "fmt"

// pass a copy on params
func invertSignalWithoutMemRef(number int) int {
	number *= -1
	return number
}

// pass a memory reference on params
func invertSignalMemRef(number *int) int {
	*number *= -1
	return *number
}

func main() {
	// before pointer
	numNotPointed := 20
	inverted := invertSignalWithoutMemRef(numNotPointed)
	fmt.Printf("Variable: %d | Result: %d\n", numNotPointed, inverted)

	// after pointer
	inverted2 := invertSignalMemRef(&numNotPointed)
	fmt.Printf("Variable: %d | Result: %d\n", numNotPointed, inverted2)
}
