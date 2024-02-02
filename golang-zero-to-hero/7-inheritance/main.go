package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	age       string
	birthdate string
}

type student struct {
	person
	course    string
	startedAt time.Time
	endedAt   time.Time
}

func main() {
	student1 := student{}
	fmt.Println(student1)
}
