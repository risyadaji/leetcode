package main

import "sort"

// Ref: https://leetcode.com/problems/maximum-length-of-pair-chain/

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	currentEnd, chainCount := -1001, 0

	for _, pair := range pairs {
		if pair[0] > currentEnd {
			currentEnd = pair[1]
			chainCount++
		}
	}

	return chainCount
}
