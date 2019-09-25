package sort

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

// BubbleSort sort itens putting each value in it's given position
func BubbleSort(v []int) {
	for changed := true; changed; {
		changed = false
		for i := 0; i < len(v)-1; i++ {
			if v[i] > v[i+1] {
				utils.Swap(v, i, i+1)
				changed = true
			}
		}
	}

}
