package main

// Ref: https://leetcode.com/problems/find-pivot-index/
// #array #prefix_sm
func pivotIndex(nums []int) int {
    n := len(nums)
    pSum := make([]int, n+1)

    for i:=1; i<=n; i++ {
        pSum[i] = nums[i-1] + pSum[i-1]
    }

    for i:=0; i<n; i++ {
        if pSum[i] == pSum[n] - pSum[i+1] {
            return i
        }
    }
    return -1
}
