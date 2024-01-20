func sumSubarrayMins(arr []int) int {
	const mod = 1e9 + 7
	var result int

	left, right := make([]int, len(arr)), make([]int, len(arr))
	stack := []int{-1}
	for i := 0; i < len(arr); i++ {
		for len(stack) > 1 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		left[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	stack = []int{len(arr)}
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 1 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		right[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	for i, n := range arr {
		result += n * (i - left[i]) * (right[i] - i)
	}

	return result % mod
}