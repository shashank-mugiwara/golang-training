package main

import (
	"fmt"
	"time"
)

func Calculation() {
	for i := 0; i < 10000; i++ {
		// some calculation
	}
}

func main() {

	start := time.Now()
	Calculation()
	end := time.Now()

	delta := end.Sub(start)
	fmt.Printf("Time taken by calculation to run and complete : %s\n", delta)
}
