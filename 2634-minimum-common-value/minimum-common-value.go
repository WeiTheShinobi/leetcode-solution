func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; ; {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
			if i == len(nums1) {
				break
			}
		} else {
			j++
			if j == len(nums2) {
				break
			}
		}
	}
	return -1
}