package main

// ref: https://leetcode.com/problems/product-of-array-except-self/?envType=study-plan-v2&envId=leetcode-75
func productExceptSelf(nums []int) []int {
	size := len(nums)   
	res := make([]int, size)
	res[0] = 1
 
	 // get prefix product multiply by related nums
	for i:=1; i<size;i++ {
		 res[i] = res[i-1] * nums[i-1]
	}
 
	 right := 1
	 // get suffix product, and save accumulate to right variable
	 for i:=size-1; i>=0; i-- {
		 res[i] *= right
		 right *= nums[i]
	 }
 
	 return res
 }


// old one
// func productExceptSelf(nums []int) []int {
// 	result := make([]int, len(nums))
// 	for i := range result {
// 		result[i] = 1
// 	}

// 	m := make(map[int]int)

// 	// O(mxn)
// 	for i := range result {
// 		// get exisisting value.
// 		if n, ok := m[nums[i]]; ok {
// 			result[i] = n
// 			continue
// 		}

// 		for j := range nums {
// 			if i != j {
// 				result[i] *= nums[j]
// 			}
// 		}

// 		// save to retrieve later
// 		m[nums[i]] = result[i]
// 	}

// 	return result
// }
