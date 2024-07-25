package main

// Ref: https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/
// #array #prefix_sum
func maxScore(s string) int {
	n := len(s)
	rSum := make([]int, n+1) // prefixSum(1)

	for i := n - 1; i > 0; i-- {
		rSum[i] = rSum[i+1]
		if s[i] == '1' {
			rSum[i]++
		}
	}

	var max, lSum int
	for i := 0; i < n; i++ {
		if s[i] == '0' && i < n-1 {
			lSum++
		}

		val := lSum + rSum[i]
		if val > max {
			max = val
		}
	}

	return max
}
