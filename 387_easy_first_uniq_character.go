package main

// Ref: https://leetcode.com/problems/first-unique-character-in-a-string/description/

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for i, char := range s {
		if total, ok := m[char]; ok && total == 1 {
			return i
		}
	}
	return -1
}
