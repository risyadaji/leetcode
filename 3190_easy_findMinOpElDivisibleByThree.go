package main

// Ref: https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/description

func minimumOperations(nums []int) int {
	// loop array
	// check is element > 3
	// if more than 3, get remaining value, if != 0 then add 1 to count (remaining will always 1 or 2)
	var n int
	for i := range nums {
		if nums[i]%3 != 0 {
			n++
		}
	}

	return n
}
