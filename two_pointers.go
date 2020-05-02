package educativeingo

import (
	"math"
	"sort"
	"strings"
)

//[1, 2, 3, 4, 6], target=6
//there is also another with hashtable approach for this
func pairWithTargerSum(arr []int, targetSum int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == targetSum {
			return []int{i, j}
		}
		if arr[i]+arr[j] > targetSum {
			j--
		}
		if arr[i]+arr[j] < targetSum {
			i++
		}
	}
	return []int{}
}

func removeDuplicates(arr []int) int {
	i := 0
	j := i + 1
	for j < len(arr) {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
		j++
	}
	return len(arr[:i+1])
}

func sorterdArraySquares(arr []int) []int {
	result := make([]int, len(arr))
	resultIdx := len(arr) - 1
	i := 0
	j := len(arr) - 1
	for i < j {
		i2 := arr[i] * arr[i]
		j2 := arr[j] * arr[j]
		if i2 < j2 {
			result[resultIdx] = j2
			j--
		} else {
			result[resultIdx] = i2
			i++
		}
		resultIdx--
	}
	return result
}

func tripletSumToZero(arr []int) [][3]int {
	var result [][3]int
	sort.Ints(arr)
	for i, targetSum := range arr {
		j := i + 1
		k := len(arr) - 1
		for j < k {
			if arr[j]+arr[k] == -targetSum {
				result = append(result, [3]int{targetSum, arr[j], arr[k]})
				j++
				k--
				for j < k && arr[j] == arr[j-1] {
					j++
				}
				for j < k && arr[k] == arr[k+1] {
					k--
				}
			}
			if arr[j]+arr[k] > -targetSum {
				k--
			}
			if arr[j]+arr[k] < -targetSum {
				j++
			}
		}
	}
	return result
}

func tripletSumCloseToTarget(arr []int, targetSum int) int {
	closestSum := math.MaxInt64
	sort.Ints(arr)
	for i, thisTarget := range arr {
		j := i + 1
		k := len(arr) - 1
		for j < k {
			res := arr[j] + arr[k] + thisTarget
			if res == targetSum {
				return 0
			}
			if res > targetSum {
				k--
			}
			if res < targetSum {
				j++
			}
			if targetSum > res {
				cloSum := targetSum - res
				if closestSum > cloSum {
					closestSum = cloSum
				}
			} else {
				cloSum := res - targetSum
				if closestSum > cloSum {
					closestSum = cloSum
				}
			}
		}
	}
	return closestSum
}

//-1, 0, 2, 3 | 3 | 2
func tripletWithSmallerSum(arr []int, target int) int {
	sort.Ints(arr)
	tripletCount := 0
	for i := range arr {
		j := i + 1
		k := len(arr) - 1
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum < target {
				tripletCount += k - j
				j++
			} else {
				k--
			}
		}
	}
	return tripletCount
}

//Input: [2, 5, 3, 10], target=30
//Output: [2], [5], [2, 5], [3], [5, 3], [10]
func subArrayProductLessThanK(arr []int, target int) [][]int {
	var results [][]int
	product := 1
	left := 0
	for right := range arr {
		product *= arr[right]
		for product >= target && left < len(arr) {
			product /= arr[left]
			left++
		}
		var result []int
		for k := right; k >= left; k-- {
			result = append([]int{arr[k]}, result...)
			results = append(results, result)
		}
	}
	return results
}

func dutchFlag(arr []int) []int {
	low := 0
	high := len(arr) - 1
	for i := 0; i <= high; {
		if arr[i] == 0 {
			arr[i], arr[low] = arr[low], arr[i]
			low++
			i++
		} else if arr[i] == 1 {
			i++
		} else {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		}
	}
	return arr
}

// Input: [4, 1, 2, -1, 1, -3], target=1
// Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
func quadrupleSumToTarget(arr []int, target int) [][]int {
	var results [][]int
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			k := j + 1
			l := len(arr) - 1
			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]
				if sum == target {
					results = append(results, []int{arr[i], arr[j], arr[k], arr[l]})
					k++
					l--
					for k < l {
						if arr[k] == arr[k-1] {
							k++
						}
						if arr[l] == arr[l+1] {
							l--
						}
					}
				}
				if sum > target {
					l--
				}
				if sum < target {
					k++
				}
			}
		}
	}
	return results
}

func backspaceCompare(str1 string, str2 string) bool {
	chars1 := strings.Split(str1, "")
	chars2 := strings.Split(str2, "")
	index1 := len(chars1) - 1
	index2 := len(chars2) - 1

	for index1 > 0 || index2 > 0 {
		i1 := findNextValidCharIndex(chars1, index1)
		i2 := findNextValidCharIndex(chars2, index2)

		if i1 < 0 && i2 < 0 {
			return true
		}
		if i1 < 0 || i2 < 0 {
			return false
		}
		if chars1[i1] != chars2[i2] {
			return false
		}
		index1 = i1 - 1
		index2 = i2 - 1
	}
	return true
}

func findNextValidCharIndex(chars []string, index int) int {
	backspaceCount := 0
	for index >= 0 {
		if chars[index] == "#" {
			backspaceCount++
		} else if backspaceCount > 0 {
			backspaceCount--
		} else {
			break
		}
		index--
	}
	return index
}

func shortestWindowSort(arr []int) int {
	low := 0
	for low < len(arr) && arr[low] <= arr[low+1] {
		low++
	}
	if low == len(arr)-1 {
		//array is already sorted
		return 0
	}
	high := len(arr) - 1
	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}

	highestInSub := math.MinInt64
	lowestInSub := math.MaxInt64
	for i := low; i <= high; i++ {
		if arr[i] > highestInSub {
			highestInSub = arr[i]
		}
		if arr[i] < lowestInSub {
			lowestInSub = arr[i]
		}
	}

	for low > 0 && arr[low-1] > lowestInSub {
		low--
	}
	for high < len(arr) && arr[high+1] < highestInSub {
		high++
	}

	return high - low + 1
}
