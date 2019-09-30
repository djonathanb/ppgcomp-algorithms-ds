package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestStackCreation(t *testing.T) {
	s := NewStack(8)

	if len(s.vet) != 8 {
		t.Errorf("%v different from 8", len(s.vet))
	}

	s = NewStack(16)

	if len(s.vet) != 16 {
		t.Errorf("%v different from 16", len(s.vet))
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack(8)
	s.Push(7)
	s.Push(9)
	s.Push(3)

	expected := []int{7, 9, 3, 0, 0, 0, 0, 0}
	if utils.Different(s.vet, expected) {
		t.Errorf("%v different from %d", s.vet, expected)
	}

	if s.tos != 2 {
		t.Errorf("%d different from (2)", s.tos)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack(8)
	s.Push(7)
	s.Push(9)
	s.Push(3)

	v, _ := s.Pop()

	if v != 3 {
		t.Errorf("%d different from (3)", v)
	}

	expected := []int{7, 9, 3, 0, 0, 0, 0, 0}
	if utils.Different(s.vet, expected) {
		t.Errorf("%v different from %d", s.vet, expected)
	}

	if s.tos != 1 {
		t.Errorf("%d different from (1)", s.tos)
	}
}
