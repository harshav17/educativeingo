package educativeingo

import "sort"

//FindSubsets of all give sets
func FindSubsets(nums []int) [][]int {
	results := [][]int{{}}
	for _, v := range nums {
		for _, v2 := range results {
			v2 = append(v2, v)
			results = append(results, v2)
		}
	}
	return results
}

//FindUniqueSubsets finds all unique subsets
func FindUniqueSubsets(nums []int) [][]int {
	sort.Ints(nums)
	results := [][]int{{}}
	startIdx, endIdx := 0, 0
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			startIdx = endIdx
		}

		endIdx = len(results)
		for i := startIdx; i < endIdx; i++ {
			v2 := results[i]
			v2 = append(v2, v)
			results = append(results, v2)
		}
	}
	return results
}

//FindPermutations finds all permutations
func FindPermutations(nums []int) [][]int {
	var results [][]int
	permutations := [][]int{{}}
	for _, v := range nums {
		permLen := len(permutations)
		for i := 0; i < permLen; i++ {
			top := permutations[0]
			permutations = permutations[1:]
			for j := 0; j <= len(top); j++ {
				newTop := top
				newTop = append(newTop[:j], append([]int{v}, newTop[j:]...)...)
				if len(newTop) == len(nums) {
					results = append(results, newTop)
				} else {
					permutations = append(permutations, newTop)
				}
			}
		}
	}
	return results
}

//FindPermutationsRecursive finds the above recursively
func FindPermutationsRecursive(nums []int) [][]int {
	var results [][]int
	return results
}
