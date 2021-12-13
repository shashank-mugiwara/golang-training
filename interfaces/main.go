package main

import "fmt"

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	len float32
	brd float32
}

// register a menthod for the interface: Shape -> Rectangle
func (rect Rectangle) Area() float32 {
	return rect.len * rect.brd
}

// register a method for the interface: Shape -> Square
func (sqr Square) Area() float32 {
	return sqr.side * sqr.side
}

func main() {

	r := Rectangle{10, 20}
	s := Square{10.3}

	shapes := []Shape{r, s}
	for _, shape := range shapes {
		area := shape.Area()
		fmt.Printf("The are of %s is %f\n", shape, area)
	}
}
