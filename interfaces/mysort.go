package mysort

type SortInterface interface {
	Len() int
	Less(i int, j int) bool
	Swap(i int, j int)
}

func IsSorted(item SortInterface) bool {
	n := item.Len()

	for i := n - 1; i > 0; i-- {
		if item.Less(i, i-1) {
			return false
		}
	}

	return true
}

// Integer type
type IntSlice []int

func (item IntSlice) Len() int {
	return len(item)
}

func (item IntSlice) Less(i int, j int) bool {
	return item[i] < item[j]
}

// String type
type StringSlice []string

func (item StringSlice) Len() int {
	return len(item)
}
func (item StringSlice) Less(i int, j int) bool {
	return item[i] < item[j]
}

func (item IntSlice) Swap(i int, j int) {
	temp := item[i]
	item[i] = item[j]
	item[j] = temp
}
