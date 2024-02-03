func largestSubmatrix(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		sort.Ints(matrix[i])
		for j := 0; j < len(matrix[0]); j++ {
			result = max(result, matrix[i][j]*(len(matrix[0])-j))
		}
	}

	return result
}