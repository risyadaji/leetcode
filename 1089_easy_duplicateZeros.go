package main

func duplicateZeros(arr []int) {
	i := 0
	for i < len(arr) {
		if arr[i] == 0 {
			shiftRight(i, arr)
			i++
		}
		i++
	}
}

func shiftRight(idx int, arr []int) {
	for i := len(arr) - 1; i > idx; i-- {
		arr[i] = arr[i-1]
	}
}
