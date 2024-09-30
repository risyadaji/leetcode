package main

// Ref: https://leetcode.com/problems/maximum-number-of-balloons/description/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxNumberOfBalloons(text string) int {
	m := make(map[rune]int)
	for _, char := range text {
		m[char]++
	}

	minCount := 10001
	minCount = min(minCount, m['b'])
	minCount = min(minCount, m['a'])
	minCount = min(minCount, m['l']/2)
	minCount = min(minCount, m['o']/2)
	minCount = min(minCount, m['n'])

	return minCount
}
