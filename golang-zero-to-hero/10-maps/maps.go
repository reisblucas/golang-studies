package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "teste",
		"age":  "teste",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	teste := map[int]string{
		1: "teste",
		2: "teste2",
	}
	fmt.Println(teste)

	teste2 := map[string]uint8{
		"number": 32,
		"score":  50,
	}
	fmt.Println(teste2)

	mapOfMaps := map[string]map[string]string{
		"user": {
			"name": "teste",
			"age":  "1",
		},
	}
	fmt.Println(mapOfMaps)
	delete(mapOfMaps, "user")
	fmt.Println(mapOfMaps)

	mapEmpty := map[string]string{}
	fmt.Println(mapEmpty)
}
