func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)

	var result int
	var f func(i int)

	set := make(map[int]int)

	f = func(i int) {
		if i > len(nums)-1 {
			return
		}

		for j := i; j < len(nums); j++ {
			if set[nums[j]-k] != 0 {
				continue
			}
			set[nums[j]]++
			result++
			f(j+1)
			set[nums[j]]--
		}
	}

	f(0)
	return result
}