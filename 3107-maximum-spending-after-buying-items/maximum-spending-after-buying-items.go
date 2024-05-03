func maxSpending(values [][]int) int64 {
	var nums []int
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			nums = append(nums, values[i][j])
		}
	}

	sort.Ints(nums)
	var result int64
	for i, d := 0, 1; i < len(nums); i, d = i+1, d+1 {
		result += int64(nums[i] * d)
	}
	return result
}