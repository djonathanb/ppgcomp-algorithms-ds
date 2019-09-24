package sort

func quickSort(v []int, i int, j int) {
	if j <= i {
		return
	}

	pivotIndex := findPivot(i, j)
	swap(v, pivotIndex, j)

	k := partition(v, i-1, j, v[j])
	swap(v, k, j)

	quickSort(v, i, k-1)
	quickSort(v, k+1, j)
}

func findPivot(i int, j int) int {
	return (i + j) / 2
}

func partition(v []int, l int, r int, pivot int) int {
	for ok := true; ok; ok = l < r {
		for grow := true; grow; grow = v[l] < pivot {
			l++
		}
		for shrink := (l < r); shrink; shrink = (l < r) && pivot < v[r] {
			r--
		}

		swap(v, l, r)
	}

	return r
}

// QuickSort division and conquer algorithm to sort an array using pivot
func QuickSort(v []int) {
	quickSort(v, 0, len(v)-1)
}
