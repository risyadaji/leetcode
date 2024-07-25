package main

func leftRightDifference(nums []int) []int {
	n := len(nums)
	lSum := make([]int, n)
	rSum := make([]int, n)

	for i := 1; i < n; i++ {
		lSum[i] = lSum[i-1] + nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		rSum[i] = rSum[i+1] + nums[i+1]
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = lSum[i] - rSum[i]
		if result[i] < 0 {
			result[i] *= -1
		}
	}

	return result
}
