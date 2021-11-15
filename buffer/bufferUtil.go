package main

import (
	"bytes"
	"fmt"
)

func main() {

	var b bytes.Buffer

	b.WriteString("SHASHANK J\n")
	b.WriteString("SHASHANK K L")
	fmt.Println(b.String())
}
