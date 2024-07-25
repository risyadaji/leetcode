package main

// Ref: https://leetcode.com/problems/find-the-highest-altitude/description/
// #prefix sum, #array

func largestAltitude(gain []int) int {
	var max, alt int

	for _, g := range gain {
		alt += g
		if alt > max {
			max = alt
		}
	}

	return max
}
