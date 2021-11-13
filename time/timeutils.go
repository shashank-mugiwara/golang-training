package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {

	t := time.Now()
	fmt.Println(t)
	fmt.Println("Time now in UTC : ", t.UTC())

	// calculating number of nano seconds in a week
	week = 7 * 24 * 60 * 60 * 1e9
	week_from_now := t.Add(week)
	fmt.Println("Week from now (by adding one week) is : ", week_from_now)

	// Get Date
	fmt.Printf("The date is : %02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

	// formatting times
	fmt.Println("In RFC1123 : ", t.Format(time.RFC1123))
	fmt.Println("In ANSIC : ", t.Format(time.ANSIC))

	// Formatting from string to date
	s := t.Format("2021 10 10")
	fmt.Println("The formatted date is : ", s)
}
