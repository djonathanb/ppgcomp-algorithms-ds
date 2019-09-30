package ds

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func TestQueueCreation(t *testing.T) {
	q := NewQueue(8)

	if len(q.vet) != 8 {
		t.Errorf("%v different from 8", len(q.vet))
	}

	q = NewQueue(16)

	if len(q.vet) != 16 {
		t.Errorf("%v different from 16", len(q.vet))
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue(8)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(3)

	expected := []int{7, 9, 3, 0, 0, 0, 0, 0}
	if utils.Different(q.vet, expected) {
		t.Errorf("%v different from %d", q.vet, expected)
	}

	if q.in != 3 {
		t.Errorf("%d different from (3)", q.out)
	}

	// Tests circular property of the queue
	q.Dequeue()
	q.Dequeue()

	q.Enqueue(11)
	q.Enqueue(13)
	q.Enqueue(15)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(4)

	expected = []int{4, 9, 3, 11, 13, 15, 1, 2}
	if utils.Different(q.vet, expected) {
		t.Errorf("%v different from %d", q.vet, expected)
	}

	if q.in != 1 {
		t.Errorf("%d different from (1)", q.out)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue(8)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(3)

	v, _ := q.Dequeue()

	if v != 7 {
		t.Errorf("%d different from (7)", v)
	}

	expected := []int{7, 9, 3, 0, 0, 0, 0, 0}
	if utils.Different(q.vet, expected) {
		t.Errorf("%v different from %d", q.vet, expected)
	}

	if q.out != 1 {
		t.Errorf("%d different from (1)", q.out)
	}

	// Tests the circular property of the queue
	q.Enqueue(11)
	q.Enqueue(13)
	q.Enqueue(15)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(5)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	v, _ = q.Dequeue()

	if v != 2 {
		t.Errorf("%d different from (2)", v)
	}

	if q.out != 0 {
		t.Errorf("%d different from (1)", q.out)
	}

	expected = []int{5, 9, 3, 11, 13, 15, 1, 2}
	if utils.Different(q.vet, expected) {
		t.Errorf("%v different from %d", q.vet, expected)
	}
}
