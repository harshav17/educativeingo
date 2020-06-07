package educativeingo

import (
	"container/heap"
	"fmt"
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
	fmt.Printf("popping %d \n", lastNode)
	return lastNode
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
		return float32((*m.max)[0])/float32(2) + float32((*m.min)[0])/float32(2)
	}
	return float32((*m.max)[0])
}
