package educativeingo

import (
	"math"
	"strings"
)

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
	smallestSize := math.MaxInt16
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

//
func noRepeatSubstring(str string) int {
	chars := strings.Split(str, "")
	distinct := make(map[string]int)
	windowStart := 0
	maxLength := 0
	for i, v := range chars {
		val, found := distinct[v]
		if found {
			if windowStart < val {
				windowStart = val
			}
		}
		//points to the element after the itself
		distinct[v] = i + 1
		if maxLength < i-windowStart+1 {
			maxLength = i - windowStart + 1
		}
	}
	return maxLength
}

//aabccb
func characterReplacement(str string, k int) int {
	chars := strings.Split(str, "")
	distinct := make(map[string]int)
	windowStart := 0
	maxLength := 0
	maxRepeatCharLength := 0
	//if windowEnd - windowStart = 3, k = 2 and distinct = 3, then no go, move on
	for i, v := range chars {
		distinct[v]++
		if maxRepeatCharLength < distinct[v] {
			maxRepeatCharLength = distinct[v]
		}
		if i-windowStart+1-maxRepeatCharLength > k {
			distinct[chars[windowStart]]--
			windowStart++
		}
		if maxLength < i-windowStart+1 {
			maxLength = i - windowStart + 1
		}
	}
	return maxLength
}

//0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1
func replacingOnes(arr []int, k int) int {
	oneCount := 0
	windowStart := 0
	maxLength := 1
	for i, v := range arr {
		if v == 1 {
			oneCount++
		}
		if i-windowStart+1-oneCount > k {
			if arr[windowStart] == 1 {
				oneCount--
			}
			windowStart++
		}
		if maxLength < i-windowStart+1 {
			maxLength = i - windowStart + 1
		}
	}
	return maxLength
}

//oidbcaf, abc
func stringPermutation(str string, pattern string) bool {
	chars := strings.Split(str, "")
	patternChars := strings.Split(pattern, "")
	patternDist := make(map[string]int)
	windowStart := 0
	matched := 0
	for _, v := range patternChars {
		patternDist[v]++
	}
	for i, v := range chars {
		_, found := patternDist[v]
		if found {
			patternDist[v]--
			if patternDist[v] == 0 {
				matched++
			}
		}

		if len(patternDist) == matched {
			return true
		}

		if i >= len(patternChars)-1 {
			c, f := patternDist[chars[windowStart]]
			if f {
				if c == 0 {
					matched--
				}
				patternDist[chars[windowStart]]++
			}
			windowStart++
		}
	}
	return false
}

//String="ppqp", Pattern="pq"
func findAnagrams(str string, pattern string) []int {
	chars := strings.Split(str, "")
	patternChars := strings.Split(pattern, "")
	patternDist := make(map[string]int)
	windowStart := 0
	matched := 0
	result := []int{}
	for _, v := range patternChars {
		patternDist[v]++
	}
	for i, v := range chars {
		_, found := patternDist[v]
		if found {
			patternDist[v]--
			if patternDist[v] == 0 {
				matched++
			}
		}

		if len(patternDist) == matched {
			result = append(result, windowStart)
		}

		if i >= len(patternChars)-1 {
			c, f := patternDist[chars[windowStart]]
			if f {
				if c == 0 {
					matched--
				}
				patternDist[chars[windowStart]]++
			}
			windowStart++
		}
	}
	return result
}

//String="aabdec", Pattern="abc"
//String="abdabca", Pattern="abc"
func minimumWindowSubstring(str string, pattern string) string {
	chars := strings.Split(str, "")
	patternChars := strings.Split(pattern, "")
	patternDist := make(map[string]int)
	windowStart := 0
	smallestLength := len(chars) + 1
	subStrStart := 0
	matched := 0
	returnStr := ""
	for _, v := range patternChars {
		patternDist[v]++
	}
	for i, v := range chars {
		_, found := patternDist[v]
		if found {
			patternDist[v]--
			if patternDist[v] >= 0 {
				matched++
			}
		}

		for len(patternDist) == matched {
			if smallestLength > i-windowStart+1 {
				smallestLength = i - windowStart + 1
				subStrStart = windowStart
			}

			c, f := patternDist[chars[windowStart]]
			if f {
				if c == 0 {
					matched--
				}
				patternDist[chars[windowStart]]++
			}
			windowStart++
		}
	}
	if smallestLength <= len(str) {
		returnStr = str[subStrStart : subStrStart+smallestLength]
	}
	return returnStr
}

//String="catfoxcat", Words=["cat", "fox"]
func wordConcatenation(str string, words []string) []int {
	wordMap := make(map[string]int)
	wordCount := len(words)
	wordLength := len(words[0])
	for _, word := range words {
		wordMap[word]++
	}
	result := []int{}
	for i := 0; i <= len(str)-wordCount*wordLength; i++ {
		wordSeen := make(map[string]int)
		for j := 0; j < wordCount; j++ {
			nextWordIndex := i + j*wordLength

			word := str[nextWordIndex : nextWordIndex+wordLength]
			val, found := wordMap[word]
			if !found {
				break
			}
			wordSeen[word]++

			if wordSeen[word] > val {
				break
			}

			if j+1 == wordCount {
				result = append(result, i)
			}
		}
	}
	return result
}
