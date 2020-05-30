package educativeingo

import (
	"container/heap"
	"sort"
)

type heapySort [][]int

func (s heapySort) Len() int {
	return len(s)
}

func (s heapySort) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s heapySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *heapySort) Push(x interface{}) {
	*s = append(*s, x.([]int))
}

func (s *heapySort) Pop() interface{} {
	old := *s
	n := len(*s)
	*s = old[0 : n-1]
	return old[n-1]
}

func mergeIntervals(intervals [][]int) [][2]int {
	sort.Sort(heapySort(intervals))
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

func conflictingAppointments(intervals [][]int) bool {
	sort.Sort(heapySort(intervals))
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			return false
		}
	}
	return true
}

func minimumMeetingRooms(intervals [][]int) int {
	sort.Sort(heapySort(intervals))
	minRooms := 0
	heapS := &heapySort{}
	for i := range intervals {
		for len(*heapS) > 0 && intervals[i][0] > (*heapS)[len(*heapS)-1][1] {
			heap.Pop(heapS)
		}
		heap.Push(heapS, intervals[i])

		if minRooms < len(*heapS) {
			minRooms = len(*heapS)
		}
	}
	return minRooms
}

func findMaxCPULoad(jobs [][]int) int {
	sort.Sort(heapySort(jobs))
	maxCPULoad := 0
	currentCPULoad := 0
	heapS := &heapySort{}
	for i := range jobs {
		for len(*heapS) > 0 && jobs[i][0] > (*heapS)[len(*heapS)-1][1] {
			j := heap.Pop(heapS).([]int)
			currentCPULoad -= j[2]
		}
		heap.Push(heapS, jobs[i])
		currentCPULoad += jobs[i][2]

		if maxCPULoad < currentCPULoad {
			maxCPULoad = currentCPULoad
		}
	}
	return maxCPULoad
}

type employeeInt struct {
	employeeNum int
	intIndex    int
	employeeHrs []int
}

type empInts []employeeInt

func (e empInts) Len() int {
	return len(e)
}

func (e empInts) Less(i, j int) bool {
	return e[i].employeeHrs[0] < e[j].employeeHrs[0]
}

func (e empInts) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *empInts) Push(x interface{}) {
	*e = append(*e, x.(employeeInt))
}

func (e *empInts) Pop() interface{} {
	n := len(*e)
	old := *e
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func findEmployeeFreeTime(hours [][][]int) [][]int {
	empsHeap := &empInts{}
	heap.Init(empsHeap)
	var result [][]int
	//push first intervals of all employees
	for i, v := range hours {
		//no compilation warnings here because of empty interface
		heap.Push(empsHeap, employeeInt{employeeNum: i, intIndex: 0, employeeHrs: v[0]})
	}

	prevInt := (*empsHeap)[len(*empsHeap)-1].employeeHrs
	for len(*empsHeap) > 0 {
		emp := heap.Pop(empsHeap).(employeeInt)

		if prevInt[1] < emp.employeeHrs[0] {
			//not overlapping
			result = append(result, []int{prevInt[1], emp.employeeHrs[0]})
			prevInt = emp.employeeHrs
		} else {
			//overlapping
			if prevInt[1] < emp.employeeHrs[1] {
				prevInt = emp.employeeHrs
			}
		}

		//add the popped employee's next interval
		if len(hours[emp.employeeNum]) > emp.intIndex+1 {
			heap.Push(empsHeap, employeeInt{employeeNum: emp.employeeNum, intIndex: emp.intIndex + 1, employeeHrs: hours[emp.employeeNum][emp.intIndex+1]})
		}
	}
	return result
}
