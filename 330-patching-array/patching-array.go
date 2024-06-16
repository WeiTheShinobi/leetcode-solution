func minPatches(nums []int, n int) int{
	s, i, result := 1, 0, 0
	for s <= n {
		if i < len(nums) && s >= nums[i] {
			s += nums[i]
			i++
		} else {
			s *= 2
			result++
		}
	}
	return result
}