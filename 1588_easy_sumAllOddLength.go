package main

// Ref: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/description/
// #array #prefix_sum #math

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0

	for i, num := range arr {
		// /2 half of full array is odd/even length. +1 is to makesure when divide by to still round up
		contribution := ((i+1)*(n-i) + 1) / 2
		sum = sum + (contribution * num)
	}

	return sum
}
