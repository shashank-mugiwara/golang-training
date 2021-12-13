package main

import "fmt"

type Identity interface {
	Set(int32)
	Get() int32
}

type Facebook struct {
	id int32
}

func (fb *Facebook) Set(id int32) {
	fb.id = id
}

func (fb Facebook) Get() int32 {
	return fb.id
}

func main() {

	var fb Facebook
	fb.Set(12)
	fmt.Println(fb.Get())
}
