package ds

import "errors"

// Queue FIFO data structure
type Queue struct {
	vet   []int
	in    int
	out   int
	count int
}

// NewQueue returns a new Queue given a determined size
func NewQueue(size int) Queue {
	return Queue{
		vet:   make([]int, size),
		in:    0,
		out:   0,
		count: 0,
	}
}

// Enqueue insert data at the end of the Queue
func (q *Queue) Enqueue(data int) {
	if q.count < len(q.vet) {
		q.count = q.count + 1
		q.vet[q.in] = data
		q.in = (q.in + 1) % len(q.vet)
	}
}

// Dequeue remove and return the item on the start of the Stack
func (q *Queue) Dequeue() (int, error) {
	if q.count > 0 {
		q.count = q.count - 1
		temp := q.vet[q.out]
		q.out = (q.out + 1) % len(q.vet)
		return temp, nil
	}
	return -1, errors.New("empty queue")
}
