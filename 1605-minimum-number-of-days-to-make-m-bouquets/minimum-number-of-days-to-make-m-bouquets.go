func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	left, right := slices.Min(bloomDay)-1, slices.Max(bloomDay)+1

	isOk := func(day int) bool {
		flower, bouquets := 0, 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= day {
				flower++
				if flower == k {
					bouquets++
					flower = 0
				}
			} else {
				flower = 0
			}
			if bouquets >= m {
				return true
			}
		}
		return false
	}
	
	for left+1 < right {
		mid := (left + right) / 2
		if isOk(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}