package main

import "fmt"

func main() {

	TwoAdder := Adder(10)
	fmt.Printf("The towAdder returns : %d and %d\n", TwoAdder(10), TwoAdder(20))

	nAdder := NAdder()
	fmt.Println(nAdder(100))
	fmt.Println(nAdder(200))
	fmt.Println(nAdder(300))

}

// Return a function using closures
func Adder(num int) func(b int) int {
	return func(b int) int {
		return num + b
	}
}

// Return a three adder function
func ThreeAdderFun(num3 int) func(num2 int) int {
	return func(num1 int) int {
		return num1 + num3
	}
}

func NAdder() func(int) int {
	var x int

	return func(delta int) int {
		x += delta
		return x
	}
}
