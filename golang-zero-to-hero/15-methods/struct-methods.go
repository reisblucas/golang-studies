package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) underAge() bool {
	return u.age > 18
}

func (u *user) birthday() {
	u.age++
}

func (u user) savingData() {
	fmt.Printf("Saving %s data on database!\n", u.name)
}

func main() {
	adam := user{"Adam", 30}
	fmt.Println(adam.underAge())

	ada := user{"Ada", 21}
	ada.birthday()
	fmt.Println(ada.age)

	ada.savingData()
	adam.savingData()
}
