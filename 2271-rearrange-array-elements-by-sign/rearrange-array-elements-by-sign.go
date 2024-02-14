func rearrangeArray(nums []int) []int {
	result := make([]int, len(nums))

	p, n := 0, 1
	for _, num := range nums {
		if num > 0 {
			result[p] = num
			p += 2
		} else {
			result[n] = num
			n += 2
		}
	}

	return result
}