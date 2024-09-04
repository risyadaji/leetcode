package main

import "log"

func removeElement(nums []int, val int) int {
	var j int
	for i := range nums {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	log.Println(nums)
	return j

	// O(m*n)
	// n := len(nums)
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == val {
	// 		for j := i; j < n; j++ {
	// 			if j == n-1 {
	// 				nums[j] = -1
	// 			} else {
	// 				nums[j] = nums[j+1]
	// 			}
	// 		}
	// 		i-- // prevent skipped value if next index is same as val
	// 		n--
	// 	}
	// }
}

func main() {
	res := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	log.Println(res)
}
