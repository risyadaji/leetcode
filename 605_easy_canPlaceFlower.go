package main

// canPlaceFlowers: https://leetcode.com/problems/can-place-flowers/
// #easy
func canPlaceFlowers(flowerbed []int, n int) bool {
	total := len(flowerbed)

	for i := 0; i < total; i++ {
		if flowerbed[i] == 0 {
			emptyLeft := (i == 0) || flowerbed[i-1] == 0
			emptyRight := (i == total-1) || flowerbed[i+1] == 0

			if emptyLeft && emptyRight {
				flowerbed[i] = 1
				n--
				i++
				if n == 0 {
					return true
				}
			}
		}
	}
	return n <= 0
}
