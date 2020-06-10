package educativeingo

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
