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
	results := [][]int{{}}
	return results
}
