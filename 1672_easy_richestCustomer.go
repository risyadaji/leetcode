package main

// ref: https://leetcode.com/problems/richest-customer-wealth/description/

func maximumWealth(accounts [][]int) int {
	var max int
	for i := range accounts {
		var sum int
		for j := range accounts[i] {
			sum += accounts[i][j]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
