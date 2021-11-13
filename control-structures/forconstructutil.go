package main

import "fmt"

func main() {

	var sum int32 = 0

	// for is the while equivalent in golang
	for sum < 10 {
		sum += 1
		fmt.Println("Sum is : ", sum)
	}

	var statement string = "Hello World!"
	for pos, char := range statement {
		fmt.Printf("The character at position %d is %c\n", pos, char)
	}
}
