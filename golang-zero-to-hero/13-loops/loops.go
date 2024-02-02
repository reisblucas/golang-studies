package main

import (
	"fmt"
	"time"
)

func main() {
	// while
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second / 10)
	}

	fmt.Println("------------")
	for j := 0; j < 10; {
		fmt.Println(j)
		j++
		time.Sleep(time.Second / 10)
	}

	fmt.Println("------------")
	// traditional loop with initialization and incrementation
	for k := 0; k < 10; k++ {
		fmt.Println(k)
		time.Sleep(time.Second / 10)
	}

	fmt.Println("------------")
	names := [3]string{"Ada", "Bob", "Claire"}
	for index, value := range names {
		fmt.Printf("index: %d | value: %s\n", index, value)
	}

	fmt.Println("------------")
	for _, value := range names {
		fmt.Printf("index: %T | value: %s\n", nil, value)
	}

	fmt.Println("-----ASCII-----")
	for i, v := range "WORD" {
		fmt.Printf("index: %d | value: %v\n", i, v)
	}

	fmt.Println("-----STRING VALUE-----")
	for i, v := range "WORD" {
		fmt.Printf("index: %d | value: %v\n", i, string(v))
	}

	fmt.Println("-----MAP-----")
	user := map[string]string{
		"name": "Ada",
		"age":  "24",
	}

	for key, value := range user {
		fmt.Printf("Key: %s | Value: %s\n", key, value)
	}

	fmt.Println("-----STRUCT-----")
	type userStruct struct {
		name string
		age  string
	}

	// for key, value := range userStruct {
	// 	fmt.Printf("Key: %s | Value: %s\n", key, value)
	// }

	fmt.Println("-----Infinite Loop-----")
	v := 10
	for {
		fmt.Printf("Running until condition be satisfied | Actual value: %d\n", v)

		if v == 15 {
			fmt.Println("Break loop value == 15")
			break
		}

		v++
		time.Sleep(time.Second)
	}
}
