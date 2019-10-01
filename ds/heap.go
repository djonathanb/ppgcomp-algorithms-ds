package ds

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// Heap Max Heap data structure implementantion
type Heap struct {
	v    []int
	size int
}

// NewHeap creates a new Heap given a determined size (capacity) or array
func NewHeap(length int) Heap {
	v := make([]int, length)
	return Heap{
		v:    v,
		size: 0,
	}
}

// NewHeapFrom creates a new Heap given values
func NewHeapFrom(v ...int) Heap {
	h := NewHeap(len(v))
	h.v = v
	h.size = len(v)
	return h
}

// Insert insert a new value (n) to the Heap
func (h *Heap) Insert(n int) {
	cur := h.size
	par := h.Parent(cur)

	h.v[cur] = n
	h.size++

	for cur > 0 && h.v[cur] > h.v[par] {
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

// GetMax return the max value (first) from heap
func (h Heap) GetMax() int {
	return h.v[0]
}

func (h Heap) maxHeapify(index int) {
	l := h.Left(index)
	r := h.Right(index)
	maior := index

	if l <= h.size-1 && h.v[l] > h.v[index] {
		maior = l
	}

	if r <= h.size-1 && h.v[r] > h.v[maior] {
		maior = r
	}

	if maior != index {
		utils.Swap(h.v, index, maior)
		h.maxHeapify(maior)
	}
}

// RemoveMax return and remove the max value from the heap and guarantee props
func (h *Heap) RemoveMax() int {
	c := h.v[0]
	h.v[0] = h.v[h.size-1]
	h.size--
	h.maxHeapify(0)
	return c
}

// BuildMaxHeap organizes a messy heap
func (h *Heap) BuildMaxHeap() {
	size := h.size - 1
	for i := size / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

// HeapSort sort the heap and return a slice of it
func (h *Heap) HeapSort() []int {
	h.BuildMaxHeap()
	s := h.size
	for i := h.size - 1; i > 0; i-- {
		utils.Swap(h.v, 0, i)
		h.size--
		h.maxHeapify(0)
	}
	return h.v[:s]
}
