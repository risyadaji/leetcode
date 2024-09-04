package main

// Ref: https://leetcode.com/problems/row-with-maximum-ones/

func rowAndMaximumOnes(mat [][]int) []int {
	maxOnesIdx := 0
	maxOnesCount := 0

	for i := 0; i < len(mat); i++ {
		onesRow := 0
		for j := 0; j < len(mat[i]); j++ {
			onesRow += mat[i][j]
		}

		if onesRow > maxOnesCount {
			maxOnesCount = onesRow
			maxOnesIdx = i
		}
	}
	return []int{maxOnesIdx, maxOnesCount}
}
