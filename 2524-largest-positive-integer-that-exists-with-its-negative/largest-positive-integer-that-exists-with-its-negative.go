func findMaxK(nums []int) int {
	result := -1
	
	m := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := m[-num]; ok {
			if num < 0 {
				num *= -1
			}
      result = max(result, num)
		}
    m[num] = struct{}{}
	}
	
	return result
}