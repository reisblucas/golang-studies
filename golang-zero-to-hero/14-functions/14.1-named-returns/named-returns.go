package main

import "fmt"

func mathCalc(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum, sub := mathCalc(1, 1)
	fmt.Println(sum, sub)
}
