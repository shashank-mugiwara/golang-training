package main

import (
	"fmt"
)

func main() {
	var slice1 []int = make([]int, 5, 5)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	newSlice1 := append(slice1, 100)
	newSlice2 := append(newSlice1, 100)
	fmt.Println("The capacity of newSlice1 is : ", cap(newSlice2))
	fmt.Println("The length of newSlice2 is : ", len(newSlice2))
}
