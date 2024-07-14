package main

// ref: https://leetcode.com/problems/running-sum-of-1d-array/description/
func runningSum(nums []int) []int {
	sum := 0
	sumNums := make([]int, len(nums))
	for i := range nums {
		sumNums[i] = nums[i] + sum
		sum += nums[i]
	}

	return sumNums
}
