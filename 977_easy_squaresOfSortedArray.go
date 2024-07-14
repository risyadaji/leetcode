package main

// ref: https://leetcode.com/problems/squares-of-a-sorted-array/description/

// Two pointer solutions
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	i := right
	result := make([]int, len(nums), len(nums))

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[i] = leftSq
			left++
		} else {
			result[i] = rightSq
			right--
		}
		i--
	}

	return result
}

// func sortedSquares(nums []int) []int {
// 	for i := range nums {
// 		nums[i] *= nums[i]
// 	}

// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] > nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 		}
// 	}

// 	return nums
// }
