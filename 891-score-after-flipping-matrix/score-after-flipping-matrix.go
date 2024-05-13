func matrixScore(grid [][]int) int {
	for i := range grid {
		if grid[i][0] == 0 {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j] ^= 1
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		num1 := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				num1++
			}
		}
		if num1 <= len(grid)/2 {
			for j := 0; j < len(grid); j++ {
				grid[j][i] ^= 1
			}
		}
	}
	var result int
	n := 1
	for i := len(grid[0]) - 1; i >= 0; i-- {
		for j := 0; j < len(grid); j++ {
			result += grid[j][i] * n
		}
		n <<= 1
	}
	return result
}