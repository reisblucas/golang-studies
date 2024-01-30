package main

import (
	"errors"
	"fmt"
)

func main() {
	number := 1000000
	fmt.Println(number)

	var integer32 int32 = 1000
	var integer64 int64 = 1000000
	fmt.Println(integer32)
	fmt.Println(integer64)

	// aliases
	var intRune rune = 12345
	fmt.Println(intRune)

	var intByte byte = 8
	fmt.Println(intByte)

	var real32 float32 = 123.45
	var real64 float64 = 1235678.70
	fmt.Println(real32, real64)

	// infer as float based on my pc archtecture - hove on variable below
	realInfered := 12345.67
	fmt.Println(realInfered)
	// end integers

	// string
	var str string = "tchurururu"
	fmt.Println(str)

	// char - position on tabela ASC
	char := 'C'
	fmt.Println(char)

	// end string

	// initial values
	var text string // equals to ""
	var num int // 0
	var err error // <nil>
	var boolean bool // false
	fmt.Println("\n", text, num, err, boolean)

	errUsage := errors.New("Intern error")
	fmt.Println(errUsage)
}