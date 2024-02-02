package main

import "fmt"

func main() {
	number := -1
	if anotherNumber := 0; anotherNumber > number {
		fmt.Printf("Another number %v is greater than number %v\n", anotherNumber, number)
	}

	str := "Teste"
	if str2 := "Teste2"; str == str2 {
		fmt.Println("Equals: " + " " + str + " " + str2)
	} else {
		fmt.Println("Not equals")
	}
}
