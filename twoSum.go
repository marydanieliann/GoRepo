package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*func main() {
	num := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
		{1, 5, 3, 8, 2},
	}

	targets := []int{9, 6, 6, 10}

	for i, nums := range num {
		result := twoSum(nums, targets[i])
		fmt.Printf("Input: %v, Target: %d, Output: %v\n", nums, targets[i], result)
	}
}
*/
