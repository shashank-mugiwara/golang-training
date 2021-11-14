package main

import "fmt"

func main() {
	Function1()
}

func Function1() {
	fmt.Println("Start of function1")

	// function2 will only be called after the completing the execution of Function2()
	defer Function2()

	fmt.Println("End of function1")

}

func Function2() {
	fmt.Println("Control entered into function2")
}
