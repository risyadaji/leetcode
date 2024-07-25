package main

import "fmt"

// Ref: https://leetcode.com/problems/ant-on-the-boundary/description/
// #array #prefix_sum

func returnToBoundaryCount(nums []int) int {
	n := len(nums)
	returnCount := 0
	pSum := make([]int, n)

	pSum[0] = nums[0]
	for i := 1; i < n; i++ {
		pSum[i] = pSum[i-1] + nums[i]
	}

	fmt.Println(pSum) // 2, 5, 0
	for _, num := range pSum {
		if num == 0 {
			returnCount++
		}
	}

	return returnCount
}
