package main

import "fmt"

func main() {
	fmt.Println("The fibonacci for 10 is : ", fibonacci(10))
	fmt.Println("The factorial of number 5 is : ", factorial(5))
}

func fibonacci(num int) (res int) {

	if num <= 1 {
		return num
	} else {
		res = fibonacci(num-1) + fibonacci(num-2)
	}

	return
}

func factorial(num uint64) (res uint64) {

	if num == 0 {
		return 1
	} else {
		res = num * factorial(num-1)
	}

	return
}
