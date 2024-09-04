package main

// Ref: https://leetcode.com/problems/matrix-diagonal-sum/

func diagonalSum(mat [][]int) int {
	totalSum := 0 // Initialize the total sum
	n := len(mat)

	j := n - 1
	for i := 0; i < n; i++ {
		totalSum += mat[i][i] // primary diagonal n == mat.length == mat[i].length means always IxI
		if j != i {
			totalSum += mat[i][j]
		}
		j--
	}
	return totalSum
}
