package problems

func RotateArray(nums []int, k int) {

	n := len(nums)

	// Reverse the entire array
	for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}

	// Reverse the first k elements
	for left, right := 0, k-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}

	// Reverse the remaining elements
	for left, right := k, n-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}

}
