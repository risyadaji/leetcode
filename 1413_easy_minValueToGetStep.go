package main

// Ref: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/description/
// #array #prefix_sum

func minStartValue(nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)

	min := 1
	for i := 1; i <= n; i++ {
		pSum[i] = pSum[i-1] + nums[i-1]
		if pSum[i] < min && pSum[i] != 0 {
			min = pSum[i]
		}
	}

	if min < 0 {
		return 1 - min
	}
	return min
}
