package main

// ref: https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/

func findNumbers(nums []int) int {
	var evenNumbers int

	for i := range nums {
		var n int
		for nums[i] > 0 {
			nums[i] /= 10
			n++
		}

		if n%2 == 0 {
			evenNumbers++
		}
	}

	return evenNumbers
}
