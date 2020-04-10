package educativeingo

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
