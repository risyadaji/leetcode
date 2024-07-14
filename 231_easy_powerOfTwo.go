package main

// Ref: https://leetcode.com/problems/power-of-two/description/?envType=daily-question&envId=2024-07-13
// tag: math, bit manipulation, recursion

func isPowerOfTwo(n int) bool {
	// 1. basic solving
	// pow := 1
	// for pow < n {
	//     pow *= 2
	// }

	// return pow == n

	// 2. using bit manipulator
	// 2: 00000010
	// 3: 00000011
	// 4: 00000100
	// 6: 00000110
	// if 3 or 6 try to AND with 3 & 2 not resulting 0, its 1
	if n == 0 {
		return false
	}

	res := (n & (n - 1)) == 0
	return res
}
