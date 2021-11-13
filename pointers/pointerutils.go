package main

import "fmt"

func main() {

	var number = 5

	var ptrToNumber = &number
	fmt.Println("The memory address of the number is : ", ptrToNumber)
	fmt.Println("The Value present at the address is : ", *ptrToNumber)

	// pointer to string
	var sampleString string = "Hello World!. I am Shashank J"
	var pointerToString *string = &sampleString
	fmt.Println("String sis stored at address : ", pointerToString)
	fmt.Println("Data present at that location is : ", *pointerToString)

	// we can now change the data at that location directly
	*pointerToString = "Data Changed."
	fmt.Println("The data at the location after the above update is : ", *pointerToString)
	fmt.Println("and the address is same : ", pointerToString)
}
