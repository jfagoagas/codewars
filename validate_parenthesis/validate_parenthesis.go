package validateparenthesis

func validateParenthesis(input string) (result bool) {
	var stack []rune
	result = true
	for _, value := range input {
		if len(stack) == 0 && value == ')' {
			result = false
			break
		}
		if value == '(' {
			// PUSH
			stack = append(stack, rune(value))
		}
		if value == ')' {
			if stack[len(stack)-1] == '(' {
				// POP
				stack = stack[:len(stack)-1]
			}
		}
	}
	return
}
