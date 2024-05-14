func getMaximumGold(grid [][]int) int {
	var f func(i, j int, seen map[int]bool) int
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	f = func(x, y int, seen map[int]bool) int {
		result := 0
		for _, d := range dir {
			nextX, nextY := x+d[0], y+d[1]
			if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
				continue
			}
			if seen[100*nextX+nextY] {
				continue
			}
			if grid[nextX][nextY] > 0 {
        seen[100*nextX+nextY] = true
				result = max(result, f(nextX, nextY, seen))
        seen[100*nextX+nextY] = false
			}
		}
		return grid[x][y] + result
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				result = max(result, f(i, j, map[int]bool{i*100+j:true}))
			}
		}
	}

	return result
}