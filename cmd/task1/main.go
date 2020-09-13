package main

import (
	"fmt"
	"github.com/RichardKhims/go_course/internal/task1"
	"github.com/RichardKhims/go_course/pkg/sort"
	"os"
)

func main() {
	sortType, arr, err := task1.ParseLineArguments(os.Args[:])
	if err != nil {
		panic("incorrect args")
	}

	res, err := sort.Sort(sortType, arr, task1.ComparatorMoreThen)

	if err != nil {
		panic(err.Error())
	}
	if res == nil {
		panic("couldn't get result")
	}

	fmt.Println(*res)
}