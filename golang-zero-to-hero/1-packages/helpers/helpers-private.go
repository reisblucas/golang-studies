package helpers

import "fmt"

// This function is an example of when we write function first letter in lowercase it refers about this function
// be private to the package

// %v - variable
// %s - string
// %d - integer
// %f - floating-point numbers
// %x - hexadecimal value
// %e - floating-point in scientific notation
// %o - octal format
// %b - binary format
// %05d - regex to limit the number of digits or characters from some value to print
// %1$d %2$s %3$.2f - numbered placeholders order
// %t - boolean value
// %T - type value print
// %c - Unicode character
// %p - pointer address
// %% - literal percent sign
func printPrivate(strPvt string) {
	fmt.Printf("Print only inside helpers package %v", strPvt)
}