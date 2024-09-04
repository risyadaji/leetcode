package main

func isValid(s string) bool {
	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	var stacks []rune
	for _, char := range s {
		n := len(stacks)
		if n > 0 {
			last := stacks[n-1]
			pair, ok := pairs[char]
			if ok && pair == last {
				stacks = stacks[:n-1]
				continue
			}
		}

		stacks = append(stacks, char)
	}

	return len(stacks) == 0
}
