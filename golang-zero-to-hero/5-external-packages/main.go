package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Printf("Testing mail validation")

	noerror := checkmail.ValidateFormat("validation.test@hotmail.com")
	fmt.Println("\nNo error:", noerror)

	error := checkmail.ValidateFormat("validation.test")
	fmt.Println("Error:", error)
}