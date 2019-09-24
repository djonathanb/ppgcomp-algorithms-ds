package sort

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// InsertionSort sort itens putting each value in it's given position
func InsertionSort(v []int) {
	for i := 0; i < len(v); i++ {
		for j := i; (j > 0) && (v[j] < v[j-1]); j-- {
			utils.Swap(v, j, j-1)
		}
	}
}
