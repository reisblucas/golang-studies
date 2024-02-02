package main

import "fmt"

type user struct {
	name string
	age  uint8
}

type address struct {
	street string
	num    uint8
}

func main() {
	fmt.Println("Structs")

	var u user
	u.name = "Ada"
	u.age = 21
	fmt.Println(u)

	user2 := user{"Ada", 23}
	fmt.Println(user2)

	user3 := user{name: "Ada"}
	fmt.Println(user3)

	address1 := address{street: "Tandandan"}
	address2 := address{street: "tundundun", num: 15}
	fmt.Println(address1)
	fmt.Println(address2)
}
