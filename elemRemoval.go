package main

func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

/*func main() {
	num := [][]int{
		{3, 2, 2, 3},
		{0, 1, 2, 2, 3, 0, 4, 2},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1},
		{},
	}

	vals := []int{3, 2, 6, 1, 1}

	for i, nums := range num {
		val := vals[i]
		k := removeElement(nums, val)
		fmt.Printf("Input: %v, Value to remove: %d, Output Length: %d, Resulting Array: %v\n", nums, val, k, nums[:k])
	}
}
*/
