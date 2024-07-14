package main

// ref: https://leetcode.com/problems/pass-the-pillow/description/?envType=daily-question&envId=2024-07-13

func passThePillow(n int, time int) int {
	i := 1
	dir := 1
	for time > 0 {
		if i == n {
			dir = -1
		} else if i == 1 {
			dir = 1
		}

		i += dir
		time--
	}

	return i
}
