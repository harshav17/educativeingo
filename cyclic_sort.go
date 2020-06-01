package educativeingo

func cyclicSort(nums []int) []int {
	for i := 0; i < len(nums); {
		v := nums[i] - 1
		if nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}
	return nums
}

func findMissingNumber(nums []int) int {
	for i := 0; i < len(nums); {
		v := nums[i]
		if v < len(nums) && nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i {
			return i
		}
	}
	return -1
}

func findAllMissingNumbers(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); {
		v := nums[i] - 1
		if nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		v := nums[i] - 1
		if nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}
	return nums[len(nums)-1]
}

func findAllDuplicates(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); {
		v := nums[i] - 1
		if nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			result = append(result, v)
		}
	}
	return result
}

func findDuplicateAndMissing(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); {
		v := nums[i] - 1
		if nums[v] != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			result = append(result, v, i+1)
		}
	}
	return result
}

func findSmallestMissingPositiveNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}

	return -1
}

func findKMissingPositiveNums(nums []int, k int) []int {
	var results []int
	xtraNos := make(map[int]bool)
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}

	for i := 0; i < len(nums) && len(results) < k; i++ {
		if nums[i] != i+1 {
			results = append(results, i+1)
			xtraNos[nums[i]] = true
		}
	}

	for i := 1; len(results) < k; i++ {
		candNo := i + len(nums)

		if !xtraNos[candNo] {
			results = append(results, candNo)
		}
	}
	return results
}
