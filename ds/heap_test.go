package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestHeapCreation(t *testing.T) {
	length := 10
	h := NewHeap(length)

	if len(h.v) != length {
		t.Errorf("vector length (%d) different from (%d)", len(h.v), length)
	}

	if h.size != 0 {
		t.Errorf("vector size (%d) different from (0)", h.size)
	}
}

func TestHeapGetMax(t *testing.T) {
	h := NewHeap(10)
	h.Insert(16)
	h.Insert(14)
	h.Insert(10)

	if h.GetMax() != 16 {
		t.Errorf("%v different from %v", h.GetMax(), 16)
	}

	h = NewHeap(10)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	if h.GetMax() != 3 {
		t.Errorf("%v different from %v", h.GetMax(), 3)
	}
}

func TestHeapParent(t *testing.T) {
	h := NewHeap(10)
	parentTest(h, 1, 0, t)
	parentTest(h, 2, 0, t)
	parentTest(h, 3, 1, t)
	parentTest(h, 4, 1, t)
	parentTest(h, 5, 2, t)
	parentTest(h, 6, 2, t)
	parentTest(h, 7, 3, t)
	parentTest(h, 8, 3, t)
	parentTest(h, 9, 4, t)
}

func parentTest(h Heap, index int, expected int, t *testing.T) {
	p := h.Parent(index)
	if p != expected {
		t.Errorf("parent (%d) different from (%d)", p, expected)
	}
}

func TestHeapLeft(t *testing.T) {
	h := NewHeap(10)
	leftTest(h, 0, 1, t)
	leftTest(h, 1, 3, t)
	leftTest(h, 2, 5, t)
	leftTest(h, 3, 7, t)
	leftTest(h, 4, 9, t)
}

func leftTest(h Heap, index int, expected int, t *testing.T) {
	l := h.Left(index)
	if l != expected {
		t.Errorf("left (%d) different from (%d)", l, expected)
	}
}

func TestHeapRight(t *testing.T) {
	h := NewHeap(10)
	rightTest(h, 0, 2, t)
	rightTest(h, 1, 4, t)
	rightTest(h, 2, 6, t)
	rightTest(h, 3, 8, t)
}

func rightTest(h Heap, index int, expected int, t *testing.T) {
	r := h.Right(index)
	if r != expected {
		t.Errorf("right (%d) different from (%d)", r, expected)
	}
}

func TestHeapInsertOrdered(t *testing.T) {
	h := NewHeap(10)
	h.Insert(16)
	h.Insert(14)
	h.Insert(10)
	h.Insert(8)
	h.Insert(7)
	h.Insert(9)
	h.Insert(3)
	h.Insert(2)
	h.Insert(4)
	h.Insert(1)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	if utils.Different(h.v, expected) {
		t.Errorf("%v different from %v", h.v, expected)
	}
}

func TestHeapInsertUnordered(t *testing.T) {
	h := NewHeap(10)
	h.Insert(12)
	h.Insert(8)
	h.Insert(3)
	h.Insert(5)
	h.Insert(9)
	h.Insert(17)
	h.Insert(6)
	h.Insert(4)
	h.Insert(15)
	h.Insert(19)

	expected := []int{19, 17, 12, 9, 15, 3, 6, 4, 5, 8}
	if utils.Different(h.v, expected) {
		t.Errorf("%v different from %v", h.v, expected)
	}
}

func TestHeapRemoveMax(t *testing.T) {
	h := NewHeapFrom(16, 14, 10, 8, 7, 9, 3, 2, 4, 1)

	// First removal
	max := h.RemoveMax()

	if max != 16 {
		t.Errorf("%v different from %v", max, 16)
	}

	current := h.Slice()
	expected := []int{14, 8, 10, 4, 7, 9, 3, 2, 1}
	if utils.Different(current, expected) {
		t.Errorf("%v different from %v", current, expected)
	}

	// Second removal
	max = h.RemoveMax()

	if max != 14 {
		t.Errorf("%v different from %v", max, 14)
	}

	current = h.Slice()
	expected = []int{10, 8, 9, 4, 7, 1, 3, 2}
	if utils.Different(current, expected) {
		t.Errorf("%v different from %v", current, expected)
	}
}

func TestHeapBuildMaxHeap(t *testing.T) {
	h := NewHeap(10)
	h.v = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	h.size = 10

	h.BuildMaxHeap()

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	if utils.Different(h.v, expected) {
		t.Errorf("%v different from %v", h.v, expected)
	}
}

func TestHeapSort(t *testing.T) {
	h := NewHeapFrom(16, 14, 10, 8, 7, 9, 3, 2, 4, 1)
	v := h.HeapSort()

	expected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	if utils.Different(v, expected) {
		t.Errorf("%v different from %v", v, expected)
	}
}
