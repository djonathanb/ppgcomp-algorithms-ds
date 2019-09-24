package sort

import "testing"

func TestInsertionSortRandom(t *testing.T) {
	numbers := []int{42, 20, 17, 13, 28, 14, 23, 15}
	InsertionSort(numbers)

	expected := []int{13, 14, 15, 17, 20, 23, 28, 42}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestInsertionSortUnordered(t *testing.T) {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3}
	InsertionSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestInsertionSortOrdered(t *testing.T) {
	numbers := []int{3, 4, 5, 6, 7, 8, 9, 10}
	InsertionSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}
