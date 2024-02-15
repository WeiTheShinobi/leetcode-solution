func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)

	sum := int64(nums[0] + nums[1])
	var result int64 = -1

	for _, n := range nums[2:] {
		if sum > int64(n) {
			result = sum + int64(n)
		}
		sum += int64(n)
	}

	return result
}