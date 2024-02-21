func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	pm := int(math.Log2(float64(m)))
	pn := int(math.Log2(float64(n)))

	if pm != pn {
		return 0
	}

	newm := 1 << pm
	return newm + rangeBitwiseAnd(m&(newm-1), n&(newm-1))
}