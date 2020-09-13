package sort

func BubbleSort(arr []int, compare func(int,int) bool) *[]int {
	x := make([]int, len(arr))
	copy(x[:], arr)

	for hasSwap := true; hasSwap;  {
		hasSwap = false
		for i := 0; i < len(x) - 1; i++ {
			if compare(x[i], x[i+1]) {
				Swap(&x[i], &x[i+1])
				hasSwap = true
			}
		}
	}
	
	return &x
}