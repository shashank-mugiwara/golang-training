package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var incorrectStringNumber string = "123.45f23"

	// absorb the erorr without capturing it
	num, _ := strconv.Atoi(incorrectStringNumber)
	fmt.Println("The converted number is : ", num)

	// capture the error
	num, err := strconv.Atoi(incorrectStringNumber)
	if err == nil {
		fmt.Println("Conversion for successful, the converted number is : ", num)
	} else {
		fmt.Println("Convertion failed.")
	}

	// switch on array of words (string)
	var statement string = "wink check nateyo kono mene ameo"
	var words []string = strings.Fields(statement)
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "check":
			fmt.Println("The word is check!")
			break
		case "wink":
			fmt.Println("The word is wink")
			break
		default:
			fmt.Println("Reached default case. Exiting")
			break
		}
	}

}
