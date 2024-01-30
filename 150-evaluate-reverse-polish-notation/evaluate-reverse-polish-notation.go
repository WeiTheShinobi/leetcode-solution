func evalRPN(tokens []string) int {
	var stack []string

	const (
		plus  = "+"
		minus = "-"
		mul   = "*"
		div   = "/"
	)

	for _, token := range tokens {
		switch token {
		case plus:
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(a+b))
		case minus:
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(a-b))
		case mul:
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(a*b))
		case div:
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(a/b))
		default:
			stack = append(stack, token)
		}
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}
