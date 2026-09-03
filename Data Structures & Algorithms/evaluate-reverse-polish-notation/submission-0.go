func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
			case "+":
				a := stack[len(stack) - 2]
				b := stack[len(stack) - 1]
				c := a + b
				stack = append(stack[:len(stack) - 2], c)
			case "-":
				a := stack[len(stack) - 2]
				b := stack[len(stack) - 1]
				c := a - b
				stack = append(stack[:len(stack) - 2], c)
			case "*":
				a := stack[len(stack) - 2]
				b := stack[len(stack) - 1]
				c := a * b
				stack = append(stack[:len(stack) - 2], c)
			case "/":
				a := stack[len(stack) - 2]
				b := stack[len(stack) - 1]
				c := a / b
				stack = append(stack[:len(stack) - 2], c)
			default:
				num, _ := strconv.Atoi(token)
				stack = append(stack, num)
		}
	}

	return stack[0]
}
