func checkSubarraySum(nums []int, k int) bool {
	m := map[int]int{0:-1}
	s := 0
	for i, n := range nums {
		s += n
		s %= k
		if prev, ok := m[s]; ok {
			if i-prev >= 2 {
				return true
			}
		} else {
			m[s] = i
		}
	}

	return false
}