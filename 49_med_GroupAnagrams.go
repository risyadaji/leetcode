package main

import "sort"

// Ref: https://leetcode.com/problems/group-anagrams/description

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := make([][]string, 0)

	for _, str := range strs {
		// sort char
		sorted := []rune(str)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})

		// check map, if not exists, append new index
		// set added index to map
		idx, ok := m[string(sorted)]
		if ok {
			res[idx] = append(res[idx], str)
		} else {
			res = append(res, []string{str})
			m[string(sorted)] = len(res) - 1
		}
	}
	return res
}
