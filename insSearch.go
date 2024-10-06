package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

/*func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 3, 5, 6}, 5},
		{[]int{1, 3, 5, 6}, 2},
		{[]int{1, 3, 5, 6}, 7},
		{[]int{1, 3, 5, 6}, 0},
		{[]int{1}, 1},
	}

	for _, testCase := range testCases {
		result := searchInsert(testCase.nums, testCase.target)
		fmt.Printf("Input: nums = %v, target = %d, Output: %d\n", testCase.nums, testCase.target, result)
	}
}*/
