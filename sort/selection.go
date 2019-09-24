package sort

// SelectionSort sort itens bringing the minimum value for given position
func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		lowIndex := i
		for j := len(v) - 1; j > i; j-- {
			if v[j] < v[lowIndex] {
				lowIndex = j
			}
		}
		swap(v, lowIndex, i)
	}
}
