package main

import "fmt"

func main() {

	// multidimensional slice
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("2D array : ", values)
	fmt.Println("First Row : ", values[0])
	fmt.Println("Second Row : ", values[1])

	// iterating the mutidimensional array
	for row := range values {
		for column := range values[0] {
			fmt.Println(values[row][column])
		}
	}
}
