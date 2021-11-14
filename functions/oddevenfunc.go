package main

import "fmt"

type flt func(int) bool

func main() {

	numbers := []int{1, 2, 3, 4, 5, 5, 6}
	var res []int = filter(numbers, IsOdd)

	fmt.Println("The Odd numbers are : ")
	for _, val := range res {
		fmt.Printf("%d\t", val)
	}
	fmt.Println()

	fmt.Println("The even numbers are : ")
	res = filter(numbers, IsEven)
	for _, val := range res {
		fmt.Printf("%d\t", val)
	}
	fmt.Println()

}

func IsEven(num int) bool {
	if num&1 == 0 {
		return true
	} else {
		return false
	}
}

func IsOdd(num int) bool {
	if num&1 == 1 {
		return true
	} else {
		return false
	}
}

func filter(numbers []int, f flt) []int {

	var res []int

	for _, val := range numbers {
		if f(val) {
			res = append(res, val)
		}
	}

	return res
}
