package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, Anonymous Functions!")
	}()

	func(str string) {
		fmt.Printf("Hello, Anonymous Functions! --%s--\n", str)
	}("teste teste teste")

	result := func(str string) string {
		return fmt.Sprintf("Hello, Anonymous Functions IN result! --%s--\n", str)
	}("teste teste teste")
	fmt.Println(result)
}
