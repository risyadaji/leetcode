package main

// Ref: https://leetcode.com/problems/subarray-sum-equals-k/description/
// #hash #prefix_sum

func subarraySum(nums []int, k int) int {
	sum := 0
	count := 0
	pSum := make(map[int]int)
	pSum[0] = 1

	for _, num := range nums {
		sum += num
		count += pSum[sum-k]
		pSum[sum]++
	}

	return count
}
