package sort

func ShakerSort(arr []int, compare func(int,int) bool) *[]int {
	x := make([]int, len(arr))
	copy(x[:], arr)
	f := func(i *int, j *int) {
		*i = *i + 1
		*j = *j - 1
	}

	for hasSwap := true; hasSwap;  {
		hasSwap = false
		for i,j := 0,len(x)-1; i < j; f(&i, &j) {
			if compare(x[i], x[i+1]) {
				Swap(&x[i], &x[i+1])
				hasSwap = true
			}
			if compare(x[j-1], x[j]) {
				Swap(&x[j-1], &x[j])
				hasSwap = true
			}
		}
	}

	return &x
}