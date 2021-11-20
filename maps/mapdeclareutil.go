package main

import (
	"fmt"
	"strconv"
)

func main() {

	var usnToNameMap map[string]int
	usnToNameMap = map[string]int{"shashank": 556761, "shashank kl": 2932893}

	fmt.Println(usnToNameMap)

	// always use make to create maps
	mp1 := make(map[string]float32)
	mp2 := make(map[string]float32)

	mp1["1"] = 334.1
	mp2["3"] = 23.1
	mp1["0"] = 322.1
	mp1["2"] = 343.1
	mp1["4"] = 999.5
	mp1["5"] = 454.2

	fmt.Println("Contents of map1 are : ", mp1)
	fmt.Println("Contents of map2 are : ", mp2)

	// slices as map values

	lol := make(map[int][]int)
	lol1 := make(map[int]*[]int)

	lol[1] = []int{10, 20, 30}

	fmt.Println("Contents of mp1 : ", lol, lol1)

	val, isPresent := mp1["shanks"]
	fmt.Println("The value is present ? ", isPresent)
	if isPresent {
		fmt.Println("The value is : ", val)
	}

	delete(mp1, "1")

	// maps of maps
	items := make([]map[string]string, 5)

	for i := range items {
		items[i] = make(map[string]string)
		items[i][strconv.Itoa(i)+"shashank"] = "Shashank J"
	}

	// print this map
	fmt.Println("The nested map structure is : ", items)

	// iterating
	for _, mapValue := range items {
		for key, value := range mapValue {
			fmt.Printf("Nested key %s for value %s\n", key, value)
		}
	}
}
