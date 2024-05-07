func minimumArrayLength(nums []int) int {
	m := slices.Min(nums)
	for _, num := range nums {
		if num%m > 0 {
			return 1
		}
	}

	count := 0
	for _, n := range nums {
		if m == n {
			count++
		}
	}
	return (count + 1) / 2
}