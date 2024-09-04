package main

// Ref: https://leetcode.com/problems/build-array-from-permutation/description/
// Great solutions: https://leetcode.com/problems/build-array-from-permutation/description/
func buildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		ans[i] = nums[num]
	}

	return ans
}

// [0,2,1,5,3,4]
// nums[nums[i]]
// [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
// [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]
// from the example we can assume nums[nums[i]] same as nums[num] -> value on eac indeces
// [0, 1, 2, 4, 5, 3]
