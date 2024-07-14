package main

// Problem: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/description/
// Solution ref: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/solutions/5459736/go-reverse-parentheses-approach-time-o-n-space-o-n

func reverseParentheses(s string) string {
	var stack []rune

	for _, ch := range s {
		if ch == ')' {
			temp := []rune{}

			// pop tail of stack and put into temp and will be added later
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1] // pop '(
			}

			stack = append(stack, temp...)
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}
