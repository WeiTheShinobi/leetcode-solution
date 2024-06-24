func minKBitFlips(nums []int, k int) int {
	var result int
	b := 0
	for i := 0; i < len(nums); i++ {
		if i-k >= 0 && nums[i-k] > 1 {
			b ^= 1
		}
		if nums[i] == b {
			if i+k > len(nums) {
				return -1
			}
			result++
			b ^= 1
			nums[i] += 2
		}
	}
	return result
}