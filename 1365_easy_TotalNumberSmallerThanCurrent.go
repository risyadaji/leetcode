package main

// Ref: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// #Array #Hash-Table #Sorting #Counting

// O(n+k) approach
func smallerNumbersThanCurrent(nums []int) []int {
	counts := [101]int{}

	// count total occurance in counts
	for _, num := range nums {
		counts[num]++
	}

	// iterate through counts to counts to calculate total smaller number from current number
	sum := 0
	for i := 0; i < len(counts); i++ {
		if counts[i] > 0 {
			// the first found number will be the smalest, so it will be 0
			counts[i], sum = sum, sum+counts[i]
		}
	}

	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		// get total smaller number
		result[i] = counts[nums[i]]
	}

	return result
}

// O(n^2) approach, not efficient
func smallerNumbersThanCurrentB(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			if nums[j] < nums[i] {
				result[i]++
			}
		}
	}

	return result
}
