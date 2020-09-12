package task1

import "strconv"

func ParseLineArguments (args []string) (string, []int, error) {
	sortType := args[1]
	x := args[2:]
	nums := make([]int, len(x))
	err := error(nil)

	for i,v := range x {
		nums[i],err = strconv.Atoi(v)
		if err != nil {
			return "", nil, err
		}
	}

	return sortType, nums, nil
}