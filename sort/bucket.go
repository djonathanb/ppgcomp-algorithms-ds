package sort

import (
	"math"

	"github.com/djonathanb/ppgcomp-algorithms-ds/ds"
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// BucketSortP (parameterized) sort itens putting each value in it's given linked list (bucket)
func BucketSortP(v []int, buckets int, maxRange int) {
	b := make([]ds.DoublyLinkedList, buckets)
	for i := 0; i < buckets; i++ {
		b[i] = ds.NewDoublyLinkedList()
	}

	for i := 0; i < len(v); i++ {
		ib := float64((v[i] - 1.0)) / (float64(maxRange) / float64(buckets))
		b[int(ib)].Insert(v[i])
	}

	i := 0
	for j := 0; j < buckets; j++ {
		list := b[j]
		// list.Print(strconv.Itoa(j) + "U")
		list.Sort()
		// list.Print(strconv.Itoa(j) + "S")
		list.Each(func(_ int, node ds.DoublyLinkedListNode) {
			v[i] = node.Data
			i++
		})
	}
}

// BucketSort sort itens putting each value in it's given linked list (bucket)
func BucketSort(v []int) {
	l := float64(len(v))
	buckets := int(math.Sqrt(l))
	max := utils.Max(v)
	BucketSortP(v, buckets, v[max])
}
