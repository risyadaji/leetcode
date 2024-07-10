package main

// ref: https://leetcode.com/problems/crawler-log-folder/description/
func minOperations(logs []string) int {
	step := 0
	for i := range logs {
		switch logs[i] {
		case "./":
            continue
		case "../":
            if step > 0 {
			    step--
            }
		default:
			step++
		}
	}

	return step
}
