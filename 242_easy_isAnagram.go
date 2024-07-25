package main

// ref: https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for i := range s {
		m[rune(s[i])]++
	}

	for i := range t {
		m[rune(t[i])]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
