// reverseVowels: https://leetcode.com/problems/reverse-vowels-of-a-string
// #easy
func reverseVowels(s string) string {
	result := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		if !isVowel(s[i]) {
			i++
			continue
		}
		if !isVowel(s[j]) {
			j--
			continue
		}
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return string(result)
}

func isVowel(c byte) bool {
	if c < 'a' {
		c += 32
	}
	return c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o'
}
