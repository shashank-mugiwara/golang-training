package main

import "fmt"

func reverse1(s string) string {
	sl := []byte(s)
	var rev [100]byte
	j := 0

	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}

	strRev := string(rev[:len(sl)])
	return strRev
}

func reverse2(s string) string {
	sl2 := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	return string(sl2)
}

// Important varient for runes (necessary for Unicode-strings!)
func reverse3(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2

	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	str := "Hello World!"
	fmt.Println("First varient of reverse : ", reverse1(str))
	fmt.Println("Second varient of reverse : ", reverse2(str))
	fmt.Println("Third Varient (for Unicode) : ", reverse3(str))
}
