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
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
		} else {
			result = append(result, [2]int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	result = append(result, [2]int{start, end})
	return result
}
