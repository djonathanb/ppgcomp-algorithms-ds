package sort

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestSelectionSortRandom(t *testing.T) {
	numbers := []int{42, 20, 17, 13, 28, 14, 23, 15}
	SelectionSort(numbers)

	expected := []int{13, 14, 15, 17, 20, 23, 28, 42}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestSelectionSortUnordered(t *testing.T) {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3}
	SelectionSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestSelectionSortOrdered(t *testing.T) {
	numbers := []int{3, 4, 5, 6, 7, 8, 9, 10}
	SelectionSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}
