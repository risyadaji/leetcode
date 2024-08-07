package main

// Ref: https://leetcode.com/problems/top-k-frequent-elements/description/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	freqArr := make([][]int, len(nums)+1)
	for num, freq := range m {
		freqArr[freq] = append(freqArr[freq], num)
	}

	res := []int{}
	for i := len(freqArr) - 1; i > 0; i-- {
		res = append(res, freqArr[i]...)
		if len(res) == k {
			return res
		}
	}
	return res
}
