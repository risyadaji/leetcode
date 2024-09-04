package main

// Ref: https://leetcode.com/problems/concatenation-of-array

func getConcatenation(nums []int) []int {
	// so basically we want to doubled up nums
	// with nums[len(n)+1] = nums[0]
	// we could utilize go "append" three dot notations as variadic
	return append(nums, nums...)
}
