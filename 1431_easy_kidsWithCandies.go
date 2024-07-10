package main

// kidsWithCandies: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
// #easy
func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	max := candies[0]
	for i := 1; i < n; i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	results := make([]bool, n)
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= max {
			results[i] = true
		}
	}

	return results
}
