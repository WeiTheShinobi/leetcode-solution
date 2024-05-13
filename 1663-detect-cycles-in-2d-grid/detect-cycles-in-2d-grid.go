func containsCycle(grid [][]byte) bool {
	type pair struct {
		x, y int
	}
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(b byte, lastX, lastY, x, y int, seen map[pair]bool) bool
	dfs = func(b byte, lastX, lastY, x, y int, seen map[pair]bool) bool {
		seen[pair{x, y}] = true
    result := false
		for _, d := range dir {
			nextX, nextY := x+d[0], y+d[1]
			if nextX == lastX && nextY == lastY {
				continue
			}
      if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
				continue
			}
			if seen[pair{nextX, nextY}] && grid[nextX][nextY] == b {
				return true
			}
			if grid[nextX][nextY] == b {
				result = dfs(b, x, y, nextX, nextY, seen)
			}
		}
		return result
	}
	
	visited := map[pair]bool{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			b := grid[i][j]
			if !visited[pair{i,j}] {
				visited[pair{i,j}] = true
				if dfs(b, -100, -100, i, j, visited) {
					return true
				}
			}
		}
	}

	return false
}