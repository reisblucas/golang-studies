package main

import "fmt"

type test struct {
	name string
}

func generic(i interface{}) {
	fmt.Println(i)
}

func main() {
	generic("tuahduhasd")
	generic(1)
	generic(true)

	fmt.Println("string", 1, 2, true, test{"Ada"})
}
