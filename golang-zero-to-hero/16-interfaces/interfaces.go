package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (f rectangle) area() float64 {
	return f.width * f.height
}

type circle struct {
	radius float64
}

func (f circle) area() float64 {
	return math.Pi * math.Pow(f.radius, 2)
}

func showArea(f form) {
	fmt.Printf("The area is %0.2f!\n", f.area())
}

func main() {
	rectangle := rectangle{width: 5, height: 10}
	showArea(rectangle)

	circle := circle{radius: 4}
	showArea(circle)
}
