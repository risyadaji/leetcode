package main

// https://leetcode.com/problems/make-the-string-great/description/
func makeGood(s string) string {
	var stack []rune

	for _, c := range s {
		// lowercase and uppercase has 32 different decimal
		if len(stack) > 0 && (stack[len(stack)-1]-c == 32 || stack[len(stack)-1]-c == -32) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
