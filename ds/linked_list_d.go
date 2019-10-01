package ds

// DoublyLinkedListNode stores data
type DoublyLinkedListNode struct {
	data int
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

// NewDoublyLinkedListNode creates a new Node storing data
func NewDoublyLinkedListNode(data int) DoublyLinkedListNode {
	return DoublyLinkedListNode{
		data: data,
		next: nil,
		prev: nil,
	}
}

// DoublyLinkedList list of nodes
type DoublyLinkedList struct {
	first *DoublyLinkedListNode
	last  *DoublyLinkedListNode
	count int
}

// NewDoublyLinkedList creates a new empty Liskend List
func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{
		first: nil,
		last:  nil,
		count: 0,
	}
}

// Insert insert a new item at the head of structure
func (l *DoublyLinkedList) Insert(data int) {
	node := NewDoublyLinkedListNode(data)

	if l.first == nil {
		l.first = &node
		l.last = &node
	} else {
		node.prev = l.last
		l.last.next = &node
		l.last = &node
	}

	l.count++
}

// Each executes func for each (index, node) in the list
func (l DoublyLinkedList) Each(fn func(int, DoublyLinkedListNode)) {
	i := 0
	node := l.first
	for node != nil {
		fn(i, *node)
		node = node.next
	}
}

// Search return a node given your key
func (l *DoublyLinkedList) Search(data int) *DoublyLinkedListNode {
	current := l.first
	for current != nil && current.data != data {
		current = current.next
	}
	return current
}

// Remove remove the given item from structure
func (l *DoublyLinkedList) Remove(data int) {
	var n *DoublyLinkedListNode = l.Search(data)

	if n == nil {
		return
	}

	isFirst := n == l.first
	isLast := n == l.last

	if isFirst && isLast {
		l.first = nil
		l.last = nil
	} else if isFirst {
		n.next.prev = nil
		l.first = n.next
	} else if isLast {
		n.prev.next = nil
		l.last = n.prev
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}

	l.count--
}

// Sort insert sort the list
func (l *DoublyLinkedList) Sort() {
	if l.count > 1 {
		current := l.first.next
		next := current.next
		for current != nil {
			node := current
			for node.prev != nil && node.data < node.prev.data {
				prev := node.prev

				if node.next != nil {
					node.next.prev = prev
				}

				if prev.prev != nil {
					prev.prev.next = next
				}

				prev.next = node.next
				node.prev = prev.prev

				prev.prev = node
				node.next = prev

				if prev == l.first {
					l.first = node
				}

				if node == l.last {
					l.last = prev
				}
			}

			current = next
			if current != nil {
				next = current.next
			} else {
				next = nil
			}
		}
	}
}

// AsArray return nodes values as array
func (l DoublyLinkedList) AsArray() []int {
	v := make([]int, l.count)
	current := l.first
	i := 0
	for current != nil {
		v[i] = current.data
		i++
		current = current.next
	}

	return v
}
