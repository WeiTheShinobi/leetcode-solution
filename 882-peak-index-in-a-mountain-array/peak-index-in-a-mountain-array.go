func peakIndexInMountainArray(arr []int) int {
	l, r := -1, len(arr)
	for l+1 != r {
		mid := (r+l)/2
		if arr[mid] < arr[mid+1] {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}