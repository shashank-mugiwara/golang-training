package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 65, 6, 7}
	res := Filter(slice, IsEven)
	fmt.Println("Higer Order Functions : The even numbers are - ", res)
}

func Filter(input []int, fn func(int) bool) []int {
	var p []int

	for _, val := range input {
		if fn(val) {
			p = append(p, val)
		}
	}

	return p
}

func IsEven(num int) bool {
	if num&1 == 1 {
		return false
	} else {
		return true
	}
}
