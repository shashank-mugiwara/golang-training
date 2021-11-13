package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// String methods
	statement := "Wink Sake Chadeo Kono Mike Rateyo"

	// strings.Contains - returns true if substr is within str
	fmt.Printf("Does statement contains Sake ? %t\n", strings.Contains(statement, "Sake"))

	// convert string to lowercase
	fmt.Printf("Converting statement to lowercase : %s\n", strings.ToLower(statement))

	// convert string to uppercase
	fmt.Printf("Converting statement to uppercase : %s\n", strings.ToUpper(statement))

	// trim off leading and trailing spaces
	stringWithSpaces := " strings with spaces "
	fmt.Printf("The statement after trimming is : \"%s\"\n", strings.TrimSpace(stringWithSpaces))

	// trim a specific part of str from the string then, the trimmable string should either be at start or end points, only
	// then the trimming occurs
	fmt.Printf("The string after trimming off the string \"Mike\" is : %s\n", strings.Trim(statement, "Wink"))

	// split the string into array of words
	var words []string = strings.Fields(statement)
	for pos, word := range words {
		fmt.Printf("Word at index %d is %s\n", pos, word)
	}

	// Join all the elements of the slice
	var joinedStringWithColon = strings.Join(words, ":")
	fmt.Printf("The statement after joining colon after each word is %s\n", joinedStringWithColon)

	// converting number to string
	fmt.Printf("The string equivalent of 33 is : %s\n", strconv.Itoa(33))

	// for converting string to number we have
	num, err := strconv.Atoi("33")
	fmt.Printf("The number converted from string to integer is : %d\nError: %s\n", num, err)

	// converting floating point to string number using FormatFloat
	fmt.Printf("Floating point converted to string is %s\n", strconv.FormatFloat(12.332, 'f', 2, 32))

	// converting string to floating point number
	floatingPointNumber, err := strconv.ParseFloat("12.3343", 32)
	if err == nil {
		fmt.Printf("The converion of string to float is successful : The number is %f\n", floatingPointNumber)
	} else {
		fmt.Printf("Parsing Failed !\n")
		fmt.Printf("Error is : %v\n", err)
	}
}
