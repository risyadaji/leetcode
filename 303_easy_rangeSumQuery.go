package main

// Ref: https://leetcode.com/problems/range-sum-query-immutable/description/
// #prefixSum #easy

type RangeQueryNumArray struct {
	sum []int
}

func NewRangeQueryNumArray(nums []int) RangeQueryNumArray {
	// O(n)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return RangeQueryNumArray{sum: nums}
}

func (this *RangeQueryNumArray) SumRange(left int, right int) int {
	// O(1)
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}

/**
[1,2,3,4,5] => i=2;j=3 => 5
[1,3,6,10,15] =>
 * Your RangeQueryNumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
*/
