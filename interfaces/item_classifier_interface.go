package main

import "fmt"

type Person struct {
	name string
	age  int
}

func classifier(items ...interface{}) {
	for index, item := range items {

		switch item.(type) {
		case bool:
			fmt.Printf("Parameter No. %d is of type boolean\n", index)
		case float32, float64:
			fmt.Printf("Parameter No. %d is of type float\n", index)
		case int, int16, int32, int64, int8:
			fmt.Printf("Parameter No. %d is of type integer\n", index)
		case nil:
			fmt.Printf("Parameter No. %d is of type nil\n", index)
		case string:
			fmt.Printf("Parameter No. %d is of type string\n", index)
		default:
			fmt.Printf("Parameter No. %d is of unknown type\n", index)
		}
	}
}

func main() {

	var integerItem interface{} = 10
	var floatItem interface{} = 10.2
	var booleanItem interface{} = false
	var stringItem interface{} = "shashank"
	var personItem interface{} = Person{name: "shashank", age: 23}

	classifier(integerItem, floatItem, booleanItem, stringItem, personItem)
}
