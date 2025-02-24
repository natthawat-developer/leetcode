package problems

func ReverseArray(arr []int) []int {
	reversed := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}

	return reversed
}
