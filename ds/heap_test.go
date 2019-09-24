package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestHeapCreation(t *testing.T) {
	length := 10
	h := New(length)

	if len(h.v) != length {
		t.Errorf("vector length (%d) different from (%d)", len(h.v), length)
	}

	if h.size != 0 {
		t.Errorf("vector size (%d) different from (0)", h.size)
	}
}

func TestParent(t *testing.T) {
	h := New(10)
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

func TestInsertOrdered(t *testing.T) {
	h := New(10)
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

func TestInsertUnordered(t *testing.T) {
	h := New(10)
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
