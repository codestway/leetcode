package easy

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	bytes := []byte(s)
	for _, b := range bytes {
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}
		switch b {
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				continue
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				continue
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, b)
	}
	return len(stack) == 0
}
