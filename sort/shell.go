package sort

import (
	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func incrInsertionSort(v []int, n int, incr int) {
	for i := incr; i < n; i += incr {
		for j := i; (j >= incr) && (v[j] < v[j-incr]); j -= incr {
			utils.Swap(v, j, j-incr)
		}
	}
}

// ShellSort insertion sort variant that pre process the array before sorting all itens
func ShellSort(v []int) {
	for i := len(v) / 2; i > 2; i /= 2 {
		for j := 0; j < i; j++ {
			incrInsertionSort(v[j:], len(v)-j, i)
		}
	}
	incrInsertionSort(v, len(v), 1)
}
