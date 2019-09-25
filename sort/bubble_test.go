package sort

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestBubbleSortRandom(t *testing.T) {
	numbers := []int{42, 20, 17, 13, 28, 14, 23, 15}
	BubbleSort(numbers)

	expected := []int{13, 14, 15, 17, 20, 23, 28, 42}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestBubbleSortUnordered(t *testing.T) {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3}
	BubbleSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}

func TestBubbleSortOrdered(t *testing.T) {
	numbers := []int{3, 4, 5, 6, 7, 8, 9, 10}
	BubbleSort(numbers)

	expected := []int{3, 4, 5, 6, 7, 8, 9, 10}

	if utils.Different(expected, numbers) {
		t.Errorf("%v different from %v", expected, numbers)
	}
}
