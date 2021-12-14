package main

import "fmt"

type AreaInterface interface {
	Area() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	side1 float32
	side2 float32
	side3 float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (square *Square) Area() float32 {
	return square.side * square.side
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.breadth * rectangle.length
}

func (triangle *Triangle) Area() float32 {
	return triangle.side1 * triangle.side2 * triangle.side3
}

func main() {
	var areaInterface AreaInterface
	square := new(Square)
	square.side = 10.4

	areaInterface = square
	fmt.Println(areaInterface.Area())
}
