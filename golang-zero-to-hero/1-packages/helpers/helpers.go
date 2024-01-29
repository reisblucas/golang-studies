package helpers

import "fmt"

// This function print something
func Printer(str string) {
	fmt.Println("Printing something", str)
	printPrivate(". Called this function inside Printer public function!")
}