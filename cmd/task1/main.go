package main

import (
	"fmt"
	"github.com/RichardKhims/go_course/internal/task1"
	"github.com/RichardKhims/go_course/pkg/sort"
	"os"
)

func main() {
	sortType, arr, err := task1.ParseLineArguments(os.Args)
	if err != nil {
		panic("Incorrect args")
	}

	var res *[]int
	comparator := func (a int, b int) bool {
		return a > b
	}
	switch sortType {
		case "bubble":
			res = sort.BubbleSort(arr[:], comparator)
			break
		case "shaker":
			res = sort.ShakerSort(arr[:], comparator)
			break
		case "selection":
			res = sort.SelectionSort(arr[:], comparator)
			break
	}

	if res == nil {
		panic("Couldn't get result")
	}

	fmt.Println(*res)
	fmt.Println(sortType)
	fmt.Println(arr)

}