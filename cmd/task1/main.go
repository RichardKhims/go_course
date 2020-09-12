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
		panic("incorrect args")
	}

	comparator := func (a int, b int) bool {
		return a > b
	}
	res, err := sort.Sort(sortType, arr, comparator)

	if err != nil {
		panic(err.Error())
	}
	if res == nil {
		panic("couldn't get result")
	}

	fmt.Println(*res)
}