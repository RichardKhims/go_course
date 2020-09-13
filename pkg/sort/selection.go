package sort

func SelectionSort(arr []int, compare func(int,int) bool) *[]int {
	x := make([]int, len(arr))
	copy(x[:], arr)

	for i := 0; i < len(x) - 1; i++ {
		cur := &x[i]
		for j := i + 1; j < len(x); j++ {
			if compare(*cur, x[j]) {
				cur = &x[j]
			}
		}
		Swap(cur, &x[i])
	}

	return &x
}