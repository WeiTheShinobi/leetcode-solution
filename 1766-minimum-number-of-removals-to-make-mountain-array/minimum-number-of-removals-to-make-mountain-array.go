func minimumMountainRemovals(nums []int) int {
	forward := make([]int, len(nums))
	back := make([]int, len(nums))

	var sl []int
	for i := 0; i < len(nums); i++ {
		j := sort.SearchInts(sl, nums[i])
		if j >= len(sl) {
			sl = append(sl, nums[i])
		} else {
			sl[j] = nums[i]
		}
		forward[i] = j + 1
	}

	sl = []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		j := sort.SearchInts(sl, nums[i])
		if j >= len(sl) {
			sl = append(sl, nums[i])
		} else {
			sl[j] = nums[i]
		}
		back[i] = j + 1
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		if forward[i] >= 2 && back[i] >= 2 {
			result = max(result, forward[i]+back[i]-1)
		}
	}

	return len(nums) - result
}