package educativeingo

import "strings"

//
func maxSubArrayOfSizeK(k int, arr []int) int {
	windowStart := 0
	maxSum := 0
	windowSum := 0
	for i, v := range arr {
		windowSum += v
		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return maxSum
}

//
func smallestSubarrayWithGivenSum(s int, arr []int) int {
	windowStart := 0
	smallestSize := 1000
	windowSum := 0
	for i, v := range arr {
		windowSum += v
		for windowSum >= s {
			if i-windowStart+1 < smallestSize {
				smallestSize = i - windowStart + 1
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return smallestSize
}

//
func longestSubsWithKDistChars(str string, k int) int {
	chars := strings.Split(str, "")
	distinct := make(map[string]int)
	windowStart := 0
	maxLength := 0
	for i, v := range chars {
		distinct[v]++
		for len(distinct) > k {
			distinct[chars[windowStart]]--
			if distinct[chars[windowStart]] == 0 {
				delete(distinct, chars[windowStart])
			}
			windowStart++
		}
		if i-windowStart+1 > maxLength {
			maxLength = i - windowStart + 1
		}
	}
	return maxLength
}
