package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestLinkedListInsert(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, 7)
	l.Insert(5, 5)
	l.Insert(3, 3)

	expected := []int{3, 5, 7}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}

func TestLinkedListSearch(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, "7")
	l.Insert(5, "5")
	l.Insert(3, "3")

	n := l.Search(7)
	if n.data != "7" {
		t.Errorf("node (%d) different of (7)", n.data)
	}

	n = l.Search(5)
	if n.data != "5" {
		t.Errorf("node (%d) different of (5)", n.data)
	}

	n = l.Search(3)
	if n.data != "3" {
		t.Errorf("node (%d) different of (3)", n.data)
	}
}

func TestLinkedListSearchAbsent(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, "7")
	l.Insert(5, "5")
	l.Insert(3, "3")

	n := l.Search(9)
	if n != nil {
		t.Errorf("node is not nil")
	}
}

func TestLinkedListRemove(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, "7")
	l.Insert(5, "5")
	l.Insert(3, "3")

	l.Remove(7)
	expected := []int{3, 5}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}

	l.Remove(3)
	expected = []int{5}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}

	l.Remove(5)
	if l.Search(5) != nil || l.head != nil {
		t.Errorf("node (5) not removed correctly")
	}
}

func TestLinkedListAbsent(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, "7")
	l.Insert(5, "5")
	l.Insert(3, "3")

	l.Remove(9)

	expected := []int{3, 5, 7}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}

func TestLinkedListEach(t *testing.T) {
	l := NewLinkedList()

	l.Insert(7, "7")
	l.Insert(5, "5")
	l.Insert(3, "3")
	l.Insert(9, "9")

	v := make([]LinkedListNode, 4)
	l.Each(func(i int, d LinkedListNode) {
		v[i] = d
	})

	expected := []int{9, 3, 5, 7}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}
