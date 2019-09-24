package ds

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
func (h Heap) Insert() {
	// current := h.size
	h.size++

}

// Parent return the index of parent node given a node index
func (h Heap) Parent(index int) int {
	return (index - 1) / 2
}
