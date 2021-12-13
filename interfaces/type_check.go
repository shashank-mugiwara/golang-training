package main

import "fmt"

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shape interface {
	Area() float32
}

func (square *Square) Area() float32 {
	return square.side * square.side
}

func (circle *Circle) Area() float32 {
	return 2 * 3.14 * circle.radius
}

func main() {

	var shape Shape
	square := Square{side: 10.22}
	shape = &square

	if _, ok := shape.(*Circle); ok {
		fmt.Printf("The shape is Circle")
	}
	if _, ok := shape.(*Square); ok {
		fmt.Println("The shape is Square")
	}
}
