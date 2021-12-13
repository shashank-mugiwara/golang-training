package main

import "fmt"

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, breadth float32
}

type Circle struct {
	radius float32
}

func (circle *Circle) Area() float32 {
	return circle.radius * 3.14 * circle.radius
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.breadth * rectangle.length
}

func (square *Square) Area() float32 {
	return square.side * square.side
}

func main() {

	square := new(Square)
	square.side = 10

	circle := new(Circle)
	circle.radius = 33.2

	rectangle := new(Rectangle)
	rectangle.breadth = 10
	rectangle.length = 20

	var shapes []Shape = []Shape{circle, square, rectangle}

	for _, shape := range shapes {
		switch t := shape.(type) {
		case *Square:
			fmt.Printf("Type is %T\n", t)
			break
		case *Rectangle:
			fmt.Printf("Type is %T\n", t)
			break
		case *Circle:
			fmt.Printf("Type is %T\n", t)
			break
		}
	}

}
