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

func TestMax(t *testing.T) {
	h := New(10)
	h.Insert(16)
	h.Insert(14)
	h.Insert(10)

	if h.Max() != 16 {
		t.Errorf("%v different from %v", h.Max(), 16)
	}

	h = New(10)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	if h.Max() != 3 {
		t.Errorf("%v different from %v", h.Max(), 3)
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

func TestLeft(t *testing.T) {
	h := New(10)
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

func TestRight(t *testing.T) {
	h := New(10)
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
