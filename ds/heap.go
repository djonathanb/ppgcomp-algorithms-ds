package ds

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// Heap Max Heap data structure implementantion
type Heap struct {
	v    []int
	size int
}

// New creates a new Heap given a determined size (capacity)
func New(length int) Heap {
	return Heap{
		v:    make([]int, length),
		size: 0,
	}
}

// Insert insert a new value to the Heap
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

// Parent return the index of parent node given a node index
func (h Heap) Parent(index int) int {
	return (index - 1) / 2
}
