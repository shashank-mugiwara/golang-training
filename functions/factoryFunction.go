package main

import (
	"fmt"
	"strings"
)

func main() {

	addBmp := MakeAddSuffix(".bmp")
	addJpg := MakeAddSuffix(".jpg")

	fmt.Println("Adding .bmp to \"2021-01-01\" : ", addBmp("2021-01-01"))
	fmt.Println("Adding .jpg to \"2021-01-01\" : ", addJpg("2021-01-01"))
}

func MakeAddSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}
