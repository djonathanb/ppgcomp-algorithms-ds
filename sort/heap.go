package sort

import "github.com/djonathanb/ppgcomp-algorithms-ds/ds"

// HeapSort sort a given array using Heap Structure
func HeapSort(v []int) []int {
	h := ds.NewHeapFrom(v...)
	return h.HeapSort()
}
