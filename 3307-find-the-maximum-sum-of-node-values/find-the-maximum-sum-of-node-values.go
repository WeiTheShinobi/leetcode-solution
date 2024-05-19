func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	f0, f1 := 0, math.MinInt
	for _, x := range nums {
		f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
	}
	return int64(f0)
}