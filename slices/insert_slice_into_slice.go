package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	insertion := []int{100, 200, 300}
	insertedSlice := insertSlice(slice, insertion, 2)

	fmt.Println("Elements after insertion into slice is : ", insertedSlice)
}

func insertSlice(slice []int, insertion []int, index int) []int {
	result := make([]int, len(slice)+len(insertion))
	numberOfElementsCopied := copy(result[0:index], slice)
	numberOfElementsCopied += copy(result[numberOfElementsCopied:], insertion)
	copy(result[numberOfElementsCopied:], slice[index:])
	return result
}
