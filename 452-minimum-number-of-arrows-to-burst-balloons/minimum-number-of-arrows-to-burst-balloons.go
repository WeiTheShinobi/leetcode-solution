func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

    result := 1
	maxRight := points[0][1]
	for _, point := range points {
		if point[0] > maxRight {
			maxRight = point[1]
			result++
		}
	}

	return result
}
