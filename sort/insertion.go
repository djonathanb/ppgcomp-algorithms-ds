package sort

// InsertionSort sort itens putting each value in it's given position
func InsertionSort(v []int) {
	for i := 0; i < len(v); i++ {
		for j := i; (j > 0) && (v[j] < v[j-1]); j-- {
			swap(v, j, j-1)
		}
	}
}
