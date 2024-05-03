func maxSpending(values [][]int) int64 {
	var nums []int
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			nums = append(nums, values[i][j])
		}
	}

	sort.Ints(nums)
	d := 1
	var result int64
	for i := 0; i < len(nums); i++ {
		result += int64(nums[i] * d)
		d++
	}
	return result
}