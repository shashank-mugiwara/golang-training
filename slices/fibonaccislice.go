package main

import "fmt"

func main() {

	resFib := fibarray(10)
	fmt.Println(resFib)
}

func fibarray(term int) []int {

	fibarr := make([]int, term)

	if term == 1 {
		fibarr[0] = 0
		return fibarr
	}

	fibarr[0], fibarr[1] = 0, 1

	for i := 2; i < term; i++ {
		fibarr[i] = fibarr[i-1] + fibarr[i-2]
	}

	return fibarr
}
