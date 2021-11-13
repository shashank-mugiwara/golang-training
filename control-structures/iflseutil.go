package main

import "fmt"

func main() {

	var num int32 = 232

	if num&1 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
