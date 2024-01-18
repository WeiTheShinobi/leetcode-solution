func maxSumBST(root *TreeNode) int {
	var (
		result int
		f      func(node *TreeNode) (int, int, int)
	)
	f = func(node *TreeNode) (int, int, int) {
		if node == nil {
      return math.MinInt, math.MaxInt, 0
		}
		lMax, lMin, lSum := f(node.Left)
		rMax, rMin, rSum := f(node.Right)

		x := node.Val
		if x <= lMax || x >= rMin {
			return math.MaxInt, math.MinInt, 0
		}
		
		s := lSum + rSum + x
		result = max(result, s)

		return max(rMax, x), min(lMin, x), s 
	}
	f(root)
	return result
}