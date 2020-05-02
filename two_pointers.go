package educativeingo

import (
	"math"
	"sort"
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
