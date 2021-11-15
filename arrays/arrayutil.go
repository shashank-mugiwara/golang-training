package main

import "fmt"

func main() {

	var arr = [5]int{3: 10, 4: 40}

	for pos, val := range arr {
		fmt.Printf("The position is : %d and value is %d\n", pos, val)
	}

	sum := ArraySum(&arr)
	fmt.Printf("The sum is : %d\n", sum)
}

func ArraySum(numbers *[5]int) int {

	var sum int = 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
