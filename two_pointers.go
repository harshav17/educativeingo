package educativeingo

import (
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
