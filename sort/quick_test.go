package sort

import "testing"

func TestQuickSortRandom(t *testing.T) {
	numbers := []int{42, 20, 17, 13, 28, 14, 23, 15}
	QuickSort(numbers)

	expected := []int{13, 14, 15, 17, 20, 23, 28, 42}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestQuickSortUnordered(t *testing.T) {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3}
	QuickSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestQuickSortOrdered(t *testing.T) {
	numbers := []int{3, 4, 5, 6, 7, 8, 9, 10}
	QuickSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}
