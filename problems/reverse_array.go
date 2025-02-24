package problems

func ReverseArray(arr []int) []int {
    left, right := 0, len(arr)-1

	for left < right {
		// สลับค่าของตำแหน่ง left และ right
		arr[left], arr[right] = arr[right], arr[left]

		// ขยับ pointer
		left++
		right--
	}

	return arr
}

// func ReverseArray(arr []int) []int {
// 	reversed := make([]int, len(arr))

// 	for i := 0; i < len(arr); i++ {
// 		reversed[i] = arr[len(arr)-1-i]
// 	}

// 	return reversed
// }