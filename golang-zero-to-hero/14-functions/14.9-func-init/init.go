package main

import "fmt"

var n int

func main() {
	fmt.Printf("n = %d\n", n)
	fmt.Println("Always after init func")
}

func init() {
	fmt.Printf("n = %d\n", n)
	fmt.Println("Always executed before main function")
	n = 10
}
