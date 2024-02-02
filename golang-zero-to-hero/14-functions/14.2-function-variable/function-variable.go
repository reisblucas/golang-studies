package main

import "fmt"

func sum(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func strAndVariaticIntegers(str string, numbers ...int) {
	for _, number := range numbers {
		fmt.Printf("%s %d\n", str, number)
	}

	if numbers == nil {
		fmt.Println(str)
	}
}

func main() {
	// 936
	sum1 := sum(1, 2, 3, 15, 20, 50, 100, -55, -200, 1000)
	fmt.Println(sum1)

	// 1000
	sum2 := sum(1, 2, 3, 15, 20, 50, 100, -55, -200, 1000, 64)
	fmt.Println(sum2)

	strAndVariaticIntegers("Uhuuu", 1, 2, 3, 4, 12, 18)
	strAndVariaticIntegers("Ohh no")
}
