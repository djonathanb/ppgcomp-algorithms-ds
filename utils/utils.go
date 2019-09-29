package utils

// Equal compares two arrays
func Equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Different compare two array
func Different(a []int, b []int) bool {
	return !Equal(a, b)
}

// Swap swaps two itens of the vector (v) given their indexes
func Swap(v []int, i int, j int) {
	copy := v[i]
	v[i] = v[j]
	v[j] = copy
}

// Max return the index of the max element of the array
func Max(v []int) int {
	if len(v) == 0 {
		return -1
	}

	max := 0
	for i := 0; i < len(v); i++ {
		if v[i] > v[max] {
			max = i
		}
	}
	return max
}
