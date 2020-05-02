package educativeingo

import "sort"

type customSort [][2]int

func (s customSort) Len() int {
	return len(s)
}

func (s customSort) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s customSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func mergeIntervals(intervals [][2]int) [][2]int {
	sort.Sort(customSort(intervals))
	var result [][2]int
	i := 1
	for i < len(intervals) {
		if intervals[i-1][1] > intervals[i][0] {
			result = append(result, [2]int{intervals[i-1][0], intervals[i][1]})
		}
	}
	return result
}
