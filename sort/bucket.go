package sort

import (
	"errors"
	"math"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// BucketSortP (parameterized) sort itens putting each value in it's given linked list (bucket)
func BucketSortP(v []int, buckets int, maxRange int) error {
	// b := make([]ds.LinkedList, buckets)
	// for i := 0; i < buckets; i++ {
	// 	b[i] = ds.NewLinkedList()
	// }

	// for i := 0; i < len(v); i++ {
	// 	ib := (v[i] - 1) / (maxRange / buckets)
	// 	b[ib].Insert(v[i])
	// }

	// for i := 0; i < len(v); {
	// 	for j := 0; j < buckets; j++ {
	// 		list := b[j]
	// 		list.Sort()
	// 		node := list.Head()
	// 		for node {
	// 			v[i] = node.data
	// 			node = node.next
	// 			i++
	// 		}
	// 	}
	// }

	return errors.New("needs doubly linked list")
}

// BucketSort sort itens putting each value in it's given linked list (bucket)
func BucketSort(v []int) error {
	l := float64(len(v))
	buckets := int(math.Sqrt(l))
	max := utils.Max(v)
	return BucketSortP(v, buckets, v[max])
}
