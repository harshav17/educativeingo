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
