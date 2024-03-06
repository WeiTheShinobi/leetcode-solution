func findLengthOfShortestSubarray(arr []int) int {
	right := len(arr) - 1

	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
  if right == 0 {
    return 0
  }

	result := right
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}

		result = min(result, right-left-1)
	}

	return result
}