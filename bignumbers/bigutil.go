package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	max64Integer := big.NewInt(math.MaxInt64)
	fmt.Println("Maximum Integer64 is : ", max64Integer)

	num1 := big.NewInt(2)
	num2 := big.NewInt(4)
	num3 := big.NewInt(7)

	mulRes := big.NewInt(1)
	mulRes.Mul(mulRes, num1).Mul(mulRes, num2).Mul(mulRes, num3)

	fmt.Println("Multiply Result : ", mulRes)
}
