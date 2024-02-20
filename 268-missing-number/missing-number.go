func missingNumber(nums []int) int {
	expSum := ((1 + len(nums)) * len(nums)) / 2
	sum := 0
	for _, n := range nums {
		sum += n
	}
	
	return expSum - sum
}