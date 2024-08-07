package main

// Ref: https://leetcode.com/problems/sort-the-people/description/
// #array #sort #hash_table

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string)

	// map name by heights
	for i, h := range heights {
		m[h] = names[i]
	}

	// sort heights
	// if need faster we can use sort.Ints lib
	heights = sortInts(heights)

	for i, h := range heights {
		names[i] = m[h]
	}

	return names
}

func sortInts(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
