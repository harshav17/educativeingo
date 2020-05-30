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

func insertIntervals(intervals [][2]int, newInterval [2]int) [][2]int {
	var result [][2]int
	i := 0
	//adds everything before the new interval
	for ; i < len(intervals) && newInterval[0] > intervals[i][1]; i++ {
		result = append(result, intervals[i])
	}

	//merge all overlapping intervals
	for ; i < len(intervals) && newInterval[1] > intervals[i][0]; i++ {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
	}
	result = append(result, newInterval)

	//add the remaining intervals
	for ; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}
	return result
}

func intervalIntersection(interval1 [][2]int, interval2 [][2]int) [][2]int {
	i := 0
	j := 0
	var result [][2]int
	for i < len(interval1) && j < len(interval2) {
		//if its overlapping
		if interval1[i][0] >= interval2[j][0] && interval1[i][0] <= interval2[j][1] || interval2[j][0] >= interval1[i][0] && interval2[j][0] <= interval1[i][1] {
			start := interval2[j][0]
			if interval1[i][0] > interval2[j][0] {
				start = interval1[i][0]
			}

			end := interval1[i][1]
			if interval1[i][1] > interval2[j][1] {
				end = interval2[j][1]
			}

			result = append(result, [2]int{start, end})
		}

		if interval1[i][1] < interval2[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}

func conflictingAppointments(intervals [][2]int) bool {
	sort.Sort(customSort(intervals))
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			return false
		}
	}
	return true
}
