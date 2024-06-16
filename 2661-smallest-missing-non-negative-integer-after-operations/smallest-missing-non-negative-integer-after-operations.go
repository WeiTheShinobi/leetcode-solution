func findSmallestInteger(nums []int, value int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] %= value
		if nums[i] < 0 {
			nums[i] += value
		}
		m[nums[i]]++
	}

	result := 0
	for {
		i := result % value
		count := m[i]
		if count == 0 {
			break
		}
		m[i]--
		result++
	}

	return result
}