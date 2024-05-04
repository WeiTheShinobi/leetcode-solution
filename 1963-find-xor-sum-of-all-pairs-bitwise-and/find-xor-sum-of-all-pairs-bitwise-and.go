func getXORSum(arr1 []int, arr2 []int) int {
	xor1 := 0
	for _, n := range arr1 {
		xor1 ^= n
	}
	xor2 := 0
	for _, n := range arr2 {
		xor2 ^=n
	}
	return xor1&xor2
}