package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

type Lener interface {
	Len(List)
}

func CountInto(a Appender, start int, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func main() {

	var list List
	CountInto(&list, 0, 100)
	fmt.Println(list)

	fmt.Println("The length of the list is :", list.Len())
}
