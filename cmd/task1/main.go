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
		os.Exit(-1)
	}

	fmt.Println(sort.BubbleSort(arr[:], func (a int, b int) bool {
		return a > b
	}))
	fmt.Println(sortType)
	fmt.Println(arr)

}