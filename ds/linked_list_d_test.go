package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestDoublyLinkedListInsert(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)

	expected := []int{7, 5, 3}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different from (%v)", l.AsArray(), expected)
	}

	if l.first.Data != 7 {
		t.Errorf("wrong start (%v)", l.first.Data)
	}

	if l.last.Data != 3 {
		t.Errorf("wrong end (%v)", l.last.Data)
	}
}

func TestDoublyLinkedListSearch(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)

	n := l.Search(7)
	if n.Data != 7 {
		t.Errorf("node (%d) different of (7)", n.Data)
	}

	n = l.Search(5)
	if n.Data != 5 {
		t.Errorf("node (%d) different of (5)", n.Data)
	}

	n = l.Search(3)
	if n.Data != 3 {
		t.Errorf("node (%d) different of (3)", n.Data)
	}
}

func TestDoublyLinkedListSearchAbsent(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)

	n := l.Search(9)
	if n != nil {
		t.Errorf("node is not nil")
	}
}

func TestDoublyLinkedListRemove(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)

	l.Remove(7)
	expected := []int{5, 3}
	if utils.Different(l.AsArray(), expected) || l.first.Data != 5 {
		t.Errorf("node (7) not removed correctly")
	}

	l.Remove(3)
	expected = []int{5}
	if utils.Different(l.AsArray(), expected) || l.last.Data != 5 {
		t.Errorf("node (3) not removed correctly")
	}

	l.Remove(5)
	if l.Search(5) != nil || l.first != nil || l.last != nil {
		t.Errorf("node (5) not removed correctly")
	}
}

func TestDoublyLinkedListAbsent(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)

	l.Remove(9)

	expected := []int{7, 5, 3}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}

func TestDoublyLinkedListEach(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Insert(7)
	l.Insert(5)
	l.Insert(3)
	l.Insert(9)

	v := make([]DoublyLinkedListNode, 4)
	l.Each(func(i int, d DoublyLinkedListNode) {
		v[i] = d
	})

	expected := []int{7, 5, 3, 9}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}

func TestDoublyLinkedListSort(t *testing.T) {
	l := NewDoublyLinkedList()
	l.Insert(7)
	l.Insert(5)
	l.Insert(3)
	l.Insert(9)
	l.Insert(1)
	l.Sort()

	expected := []int{1, 3, 5, 7, 9}
	if utils.Different(l.AsArray(), expected) {
		t.Errorf("(%v) different of (%v)", l.AsArray(), expected)
	}
}
