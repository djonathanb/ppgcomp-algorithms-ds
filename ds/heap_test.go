package ds

import "testing"

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
