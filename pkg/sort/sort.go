package sort

import "errors"

type sorter func([]int, func(int,int) bool) *[]int
var sortMap = map[string]sorter {
	"bubble" : BubbleSort,
	"shaker" : ShakerSort,
	"selection" : SelectionSort,
}

func Sort(method string, arr []int, compare func(int,int) bool) (*[]int, error) {
	sorterFunc := sortMap[method]
	if sorterFunc == nil {
		return nil, errors.New("method not found")
	}

	return sorterFunc(arr, compare), nil
}