package main

// https://leetcode.com/problems/removing-stars-from-a-string/description/
func removeStars(s string) string {
	var stack []rune

	for _, c := range s {
		if len(stack) > 0 && c == '*' {
			stack = stack[:len(stack)-1]
		} else if c != '*' {
			stack = append(stack, c)
		}
	}

	return string(stack)
}
