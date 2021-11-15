package main

import "fmt"

func main() {

	// syntax of copy function
	// copy(dest, src []type) int

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Println("Copied number of elements : ", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println("Appended sl3 elements, overall elements in total : ", sl3)
}
