package main

import "fmt"

func main() {
	fmt.Println("The sum of three numbers are : ", addThreeNmbers(getThreeNumbers(10, 20, 30)))

	variadicFunction(10, "shashank j", "shashank k l", "antman")
}

func getThreeNumbers(n1, n2, n3 int) (int, int, int) {
	return n1, n2, n3
}

func addThreeNmbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func variadicFunction(firstParam int, params ...string) {
	fmt.Println("The first parameter is : ", firstParam)
	fmt.Println("The next set of parameters are : ")

	for _, word := range params {
		fmt.Println(word)
	}
}
