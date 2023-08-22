func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := math.MaxInt32, math.MaxInt32
	for i := range nums {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
