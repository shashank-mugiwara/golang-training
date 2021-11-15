package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println("The length before enlarging is : ", len(s))
	fmt.Println(s)

	fmt.Println("Enlarging / Magnifying the slice with a factor of 3")
	s = enlarge(s, 3)
	fmt.Println("The length after enlarging is : ", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	return ns
}
