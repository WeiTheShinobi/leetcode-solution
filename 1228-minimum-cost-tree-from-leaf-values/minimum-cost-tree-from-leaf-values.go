func mctFromLeafValues(arr []int) int {
	var stack []int
	result := 0
	for _, x := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= x {
			if len(stack) == 1 || stack[len(stack)-2] > x {
				result += stack[len(stack)-1] * x
			} else {
				result += stack[len(stack)-1] * stack[len(stack)-2]
			}
			// [x, y] z  | y<z && x<z -> x*y
			// [x, z] 
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, x)
	}
	for len(stack) > 1 {
		result += stack[len(stack)-1]*stack[len(stack)-2]
		stack = stack[:len(stack)-1]
	}

	return result
}