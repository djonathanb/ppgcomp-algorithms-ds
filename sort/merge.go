package sort

func mergeSort(v []int, temp []int, left int, right int) {
	if left == right {
		return
	}

	// List of one element
	mid := (left + right) / 2
	mergeSort(v, temp, left, mid)
	mergeSort(v, temp, mid+1, right)

	// Merge sublists
	merge(v, temp, left, mid, right)
}

func merge(v []int, temp []int, left int, mid int, right int) {
	// Copy subarray to temp
	for i := left; i <= right; i++ {
		temp[i] = v[i]
	}

	// Do the merge operation back to target array
	i1 := left
	i2 := mid + 1
	for curr := left; curr <= right; curr++ {
		if i1 == mid+1 { // Left sublist exhausted
			v[curr] = temp[i2]
			i2++
		} else if i2 > right { // Right sublist exhausted
			v[curr] = temp[i1]
			i1++
		} else if temp[i1] < temp[i2] {
			v[curr] = temp[i1]
			i1++
		} else {
			v[curr] = temp[i2]
			i2++
		}
	}
}

// MergeSort division and conquer algorithm to sort an array merging subarrays
func MergeSort(v []int) {
	var temp = make([]int, len(v))
	mergeSort(v, temp, 0, len(v)-1)
}
