package min

type Miner interface {
	Len() int
	ElemAtX(x int) interface{}
	Less(i int, j int) bool
}

type IntArray []int

func (item IntArray) Len() int {
	return len(item)
}

func (item IntArray) ElemAtX(x int) int {
	return item[x]
}

func (item IntArray) Less(i int, j int) bool {
	return item[i] < item[j]
}

func Min(data Miner) interface{} {

	min := data.ElemAtX(0)
	min_idx := 0

	for i := 0; i < data.Len(); i++ {
		if data.Less(i, min_idx) {
			min = data.ElemAtX(i)
			min_idx = i
		}
	}

	return min
}
