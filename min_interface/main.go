package main

import (
	"fmt"
	"golang-training/min_interface/min"
)

func ints() {
	item := []int{10, 20, 30, 40, 50, 60, 70}
	a := min.IntArray(item)
	m := min.Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func main() {
	ints()
}
