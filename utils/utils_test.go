package utils

import "testing"

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if !Equal(a, b) {
		t.Errorf("assertion error")
	}
}

func TestNotEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 1, 3, 4}
	if Equal(a, b) {
		t.Errorf("assertion error")
	}
}

func TestDifferent(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 1, 3, 4}
	if !Different(a, b) {
		t.Errorf("assertion error")
	}
}

func TestNotDifferent(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if Different(a, b) {
		t.Errorf("assertion error")
	}
}

func TestSwap(t *testing.T) {
	a := []int{1, 2, 3, 4}
	Swap(a, 0, 2)
	expected := []int{3, 2, 1, 4}
	if Different(a, expected) {
		t.Errorf("%v different from %v", a, expected)
	}

	Swap(a, 0, 3)
	expected = []int{4, 2, 1, 3}
	if Different(a, expected) {
		t.Errorf("%v different from %v", a, expected)
	}
}

func TestMax(t *testing.T) {
	v := []int{12, 2, 3, 9, 55, -0, 54, 7}
	max := Max(v)

	if max != 4 {
		t.Errorf("(%d) different from (4)", max)
	}
}
