package main

import (
	"container/list"
	"fmt"
)

func main() {

	var simpleDoublyList list.List
	simpleDoublyList.PushBack(10)
	simpleDoublyList.PushBack("shashank")

	fmt.Println("The contents of the list are : ", simpleDoublyList)
}
