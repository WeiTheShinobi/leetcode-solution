func distributeCoins(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        d := dfs(node.Left) + dfs(node.Right) + node.Val - 1
        ans += abs(d)
        return d
    }
    dfs(root)
    return
}

func abs(x int) int { if x < 0 { return -x }; return x }