package sort

func equal(a []int, b []int) bool {
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

func different(a []int, b []int) bool {
	return !equal(a, b)
}

func swap(v []int, i int, j int) {
	copy := v[i]
	v[i] = v[j]
	v[j] = copy
}
