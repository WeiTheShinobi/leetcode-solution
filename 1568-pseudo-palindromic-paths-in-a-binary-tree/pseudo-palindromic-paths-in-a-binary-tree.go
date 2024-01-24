func pseudoPalindromicPaths(root *TreeNode) int {
	var f func(node *TreeNode, val int) int
	f = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}
		val ^= 1 << node.Val
		if node.Left == nil && node.Right == nil && bits.OnesCount(uint(val)) <= 1 {
			return 1
		}
		return f(node.Left, val) + f(node.Right, val)
	}
	return f(root, 0)
}
