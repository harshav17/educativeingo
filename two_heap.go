package educativeingo

import (
	"container/heap"
)

type minHeap []int

func (n minHeap) Len() int {
	return len(n)
}

func (n minHeap) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n minHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *minHeap) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *minHeap) Pop() interface{} {
	heapSize := len(*n)
	lastNode := (*n)[heapSize-1]
	*n = (*n)[:heapSize-1]
	return lastNode
}

//Find s a number from the heap
func (n *minHeap) Find(num int) int {
	for i, v := range *n {
		if v == num {
			return i
		}
	}
	return -1
}

type maxHeap []int

func (s maxHeap) Len() int {
	return len(s)
}

func (s maxHeap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s maxHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *maxHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *maxHeap) Pop() interface{} {
	heapSize := len(*s)
	lastNode := (*s)[heapSize-1]
	*s = (*s)[:heapSize-1]
	return lastNode
}

//Find s a number from the heap
func (s *maxHeap) Find(num int) int {
	for i, v := range *s {
		if v == num {
			return i
		}
	}
	return -1
}

//MedianOfStream gives a median of stream
type MedianOfStream struct {
	min *minHeap
	max *maxHeap
}

//Constructor median of stream
func Constructor() MedianOfStream {
	max, min := &maxHeap{}, &minHeap{}
	heap.Init(max)
	heap.Init(min)

	return MedianOfStream{
		max: max,
		min: min,
	}
}

//InsertNum inserts a number into stream
func (m *MedianOfStream) InsertNum(num int) {
	if m.max.Len() == 0 || (*m.max)[0] >= num {
		heap.Push(m.max, num)
	} else {
		heap.Push(m.min, num)
	}

	if m.max.Len() > m.min.Len()+1 {
		heap.Push(m.min, heap.Pop(m.max))
	} else if m.max.Len() < m.min.Len() {
		heap.Push(m.max, heap.Pop(m.min))
	}
}

//FindMedian finds a median of stream of numbers
func (m *MedianOfStream) FindMedian() float32 {
	if m.max.Len() == m.min.Len() {
		return float32((*m.max)[0]+(*m.min)[0]) / float32(2)
	}
	return float32((*m.max)[0])
}

//FindSlidingWindowMedian finds a
func FindSlidingWindowMedian(nums []int, k int) []float32 {
	m := Constructor()
	var results []float32
	for i, v := range nums {
		if m.max.Len() == 0 || (*m.max)[0] >= v {
			heap.Push(m.max, v)
		} else {
			heap.Push(m.min, v)
		}
		rebalanceHeaps(&m)

		if i-k+1 >= 0 {
			if m.max.Len() == m.min.Len() {
				results = append(results, float32((*m.max)[0]+(*m.min)[0])/float32(2))
			} else {
				results = append(results, float32((*m.max)[0]))
			}

			elemToBeRem := nums[i-k+1]
			if elemToBeRem <= (*m.max)[0] {
				//remove it from max heap
				elemIndex := m.max.Find(elemToBeRem)
				heap.Remove(m.max, elemIndex)
			} else {
				//remove it from min heap
				elemIndex := m.min.Find(elemToBeRem)
				heap.Remove(m.min, elemIndex)
			}
			rebalanceHeaps(&m)
		}
	}
	return results
}

func rebalanceHeaps(m *MedianOfStream) {
	if m.max.Len() > m.min.Len()+1 {
		heap.Push(m.min, heap.Pop(m.max))
	} else if m.max.Len() < m.min.Len() {
		heap.Push(m.max, heap.Pop(m.min))
	}
}

type cap struct {
	cap   int
	index int
}

type capss []cap

func (c capss) Less(i, j int) bool {
	return c[i].cap < c[j].cap
}

func (c capss) Len() int {
	return len(c)
}

func (c capss) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *capss) Pop() interface{} {
	heapSize := len(*c)
	lastNode := (*c)[heapSize-1]
	*c = (*c)[:heapSize-1]
	return lastNode
}

func (c *capss) Push(x interface{}) {
	*c = append(*c, x.(cap))
}

//profits
type prof struct {
	prof  int
	index int
}

type profs []prof

func (p profs) Less(i, j int) bool {
	return p[i].prof > p[j].prof
}

func (p profs) Len() int {
	return len(p)
}

func (p profs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *profs) Push(x interface{}) {
	*p = append(*p, x.(prof))
}

func (p *profs) Pop() interface{} {
	heapSize := len(*p)
	lastNode := (*p)[heapSize-1]
	*p = (*p)[:heapSize-1]
	return lastNode
}

//FindMaxProfit finds max profit
func FindMaxProfit(caps []int, profits []int, initCap int, numProjs int) int {
	min := &capss{}
	heap.Init(min)
	for i, v := range caps {
		heap.Push(min, cap{cap: v, index: i})
	}

	avCap := initCap

	for i := 0; i < numProjs; i++ {
		max := &profs{}
		heap.Init(max)
		for len(*min) > 0 && (*min)[0].cap <= avCap {
			top := heap.Pop(min).(cap)
			heap.Push(max, prof{prof: profits[top.index], index: top.index})
		}

		if len(*max) == 0 {
			break
		}

		best := heap.Pop(max).(prof)
		avCap += best.prof
	}

	return avCap
}

//FindNextInterval find all the next intervals
func FindNextInterval(intervals [][]int) []int {
	results := make([]int, len(intervals))
	startMax := &profs{}
	heap.Init(startMax)

	endMax := &profs{}
	heap.Init(endMax)

	for i, v := range intervals {
		heap.Push(startMax, prof{prof: v[0], index: i})
		heap.Push(endMax, prof{prof: v[1], index: i})
	}

	for i := 0; i < len(intervals); i++ {
		topEnd := heap.Pop(endMax).(prof)
		results[topEnd.index] = -1
		if (*startMax)[0].prof >= topEnd.prof {
			topStart := heap.Pop(startMax).(prof)

			for len(*startMax) > 0 && (*startMax)[0].prof >= topEnd.prof {
				topStart = heap.Pop(startMax).(prof)
			}
			results[topEnd.index] = topStart.index
			heap.Push(startMax, topStart)
		}
	}
	return results
}
