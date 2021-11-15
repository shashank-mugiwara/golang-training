package main

import (
	"fmt"
)

func main() {

	slice1 := make([]int, 0, 10)

	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : len(slice1)+1]
		slice1[i] = i
	}

	fmt.Printf("The length of the slice is : %d\n", len(slice1))
}
