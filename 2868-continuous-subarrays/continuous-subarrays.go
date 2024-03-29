func continuousSubarrays(nums []int) int64 {
	var result int64
	left := 0
	count := map[int]int{}
	for right, num := range nums {
		count[num]++
		for {
			mx, mi := num, num
			for k := range count {
				mx = max(mx, k)
				mi = min(mi, k)
			}
			if mx-mi <= 2 {
				break
			}
			y := nums[left]
			count[y]--
			if count[y] == 0 {
				delete(count, y)
			}
			left++
		}
		result += int64(right - left + 1)
	}

	return result
}