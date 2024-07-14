package main

// ref: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/description/
func numberOfSteps(num int) int {
	steps := 0

	// using bitwise
	for num != 0 {
		if (num & 1) == 0 { // xxxxxx0 & xxxxxxx1
			num >>= 1 // right shift mean divided by 2
		} else {
			num--
		}
	}

	return steps

	// using modulo
	// for num != 0 {
	// 	if num%2 == 0 {
	// 		num /= 2
	// 	} else {
	// 		num -= 1
	// 	}
	// 	ans++
	// }
	// return steps
}
