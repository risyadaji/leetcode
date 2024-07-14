package main

func balancedStringSplit(s string) int {
	var n, bal int

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			bal++
		}
		if s[i] == 'R' {
			bal--
		}
		if bal == 0 {
			n++
		}
	}

	return n
}
