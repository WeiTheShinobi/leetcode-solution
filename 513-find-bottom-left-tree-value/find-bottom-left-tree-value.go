func findBottomLeftValue(root *TreeNode) int {
	maxDepth := 0
	result := 0
	var f func(node *TreeNode, d int)
	f = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if d > maxDepth {
			maxDepth = d
			result = node.Val
		}
		if node.Left != nil {
			f(node.Left, d+1)
		}
		if node.Right != nil {
			f(node.Right, d+1)
		}
	}

	f(root, 1)
	return result
}