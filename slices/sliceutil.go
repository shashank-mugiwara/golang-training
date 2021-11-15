package main

import "fmt"

func main() {

	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	// load the array with integers
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d\t", slice1[i])
	}

	fmt.Println()

	fmt.Println("The length of arr1 is : ", len(arr1))
	fmt.Println("The length of slice1 is : ", len(slice1))
	fmt.Println("The capacity of slice1 is : ", cap(slice1))

	// grow the slice:
	slice1 = slice1[0:4]
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	// }
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// creating slice with make
	var sliceWithMake []int = make([]int, 10)
	fmt.Println("Slice created with make function, The length is : ", len(sliceWithMake))
	fmt.Println("Slice created with make function, The capacity is : ", cap(sliceWithMake))
}
