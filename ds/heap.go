package ds

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// Heap Max Heap data structure implementantion
type Heap struct {
	v         []int
	size      int
	isMaxHeap bool
	compare   func([]int, int, int) bool
}

// newHeap creates a new Heap given a determined size (capacity) or array and type (min/max)
func newHeap(length int, isMaxHeap bool) Heap {
	v := make([]int, length)

	var fn func([]int, int, int) bool
	if isMaxHeap {
		fn = func(v []int, i int, j int) bool {
			return v[i] > v[j]
		}
	} else {
		fn = func(v []int, i int, j int) bool {
			return v[i] < v[j]
		}
	}

	return Heap{
		v:         v,
		size:      0,
		isMaxHeap: isMaxHeap,
		compare:   fn,
	}
}

// newHeapFrom creates a new Heap given values
func newHeapFrom(isMaxHeap bool, v []int) Heap {
	h := newHeap(len(v), isMaxHeap)
	h.v = v
	h.size = len(v)
	return h
}

// NewMaxHeapFrom creates a new Heap given values
func NewMaxHeapFrom(v ...int) Heap {
	return newHeapFrom(true, v)
}

// NewMinHeapFrom creates a new Heap given values
func NewMinHeapFrom(v ...int) Heap {
	return newHeapFrom(false, v)
}

// NewMaxHeap creates a max heap
func NewMaxHeap(length int) Heap {
	return newHeap(length, true)
}

// NewMinHeap creates a min heap
func NewMinHeap(length int) Heap {
	return newHeap(length, false)
}

// Len returns the lenght of the heap
func (h *Heap) Len() int {
	return h.size
}

// Insert insert a new value (n) to the Heap
func (h *Heap) Insert(n int) {
	cur := h.size
	par := h.Parent(cur)

	h.v[cur] = n
	h.size++

	for cur > 0 && h.compare(h.v, cur, par) {
		utils.Swap(h.v, cur, par)
		cur = par
		par = h.Parent(cur)
	}
}

// Slice return a sliced version of the inner vector with the proper size
func (h Heap) Slice() []int {
	return h.v[:h.size]
}

// Parent return the index of parent node given a node index
func (h Heap) Parent(index int) int {
	return (index - 1) / 2
}

// Left return the index of the left node given a node index
func (h Heap) Left(index int) int {
	return (index * 2) + 1
}

// Right return the index of the right node given a node index
func (h Heap) Right(index int) int {
	return (index + 1) * 2
}

// GetTop return the top value from heap
func (h Heap) GetTop() int {
	return h.v[0]
}

func (h Heap) heapify(index int) {
	l := h.Left(index)
	r := h.Right(index)
	selected := index

	if l <= h.size-1 && h.compare(h.v, l, index) {
		selected = l
	}

	if r <= h.size-1 && h.compare(h.v, r, selected) {
		selected = r
	}

	if selected != index {
		utils.Swap(h.v, index, selected)
		h.heapify(selected)
	}
}

// RemoveTop return and remove the top value from the heap and guarantee props
func (h *Heap) RemoveTop() int {
	c := h.v[0]
	h.v[0] = h.v[h.size-1]
	h.size--
	h.heapify(0)
	return c
}

// BuildHeap organizes a messy heap
func (h *Heap) BuildHeap() {
	size := h.size - 1
	for i := size / 2; i >= 0; i-- {
		h.heapify(i)
	}
}

// HeapSort sort the max heap and return a slice of it
func (h *Heap) HeapSort() []int {
	h.BuildHeap()
	s := h.size
	if h.isMaxHeap {
		for i := h.size - 1; i > 0; i-- {
			utils.Swap(h.v, 0, i)
			h.size--
			h.heapify(0)
		}
	} else {
		for i := 0; i < h.size-1; i++ {
			utils.Swap(h.v, 0, i)
			h.size--
			h.heapify(0)
		}
	}
	return h.v[:s]
}
