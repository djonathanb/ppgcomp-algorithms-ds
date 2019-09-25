package ds

// LinkedListNode stores data
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

// LinkedList list of nodes
type LinkedList struct {
	head  *LinkedListNode
	count int
}

// NewLinkedListNode creates a new Node storing data
func NewLinkedListNode(data int) LinkedListNode {
	return LinkedListNode{
		data: data,
		next: nil,
	}
}

// NewLinkedList creates a new empty Liskend List
func NewLinkedList() LinkedList {
	return LinkedList{
		head:  nil,
		count: 0,
	}
}

// Insert insert a new item at the head of structure
func (l *LinkedList) Insert(data int) {
	node := NewLinkedListNode(data)
	node.next = l.head
	l.head = &node
	l.count++
}

// Search return a node given your key
func (l *LinkedList) Search(data int) *LinkedListNode {
	current := l.head
	for current != nil && current.data != data {
		current = current.next
	}
	return current
}

// Remove remove the given item from structure
func (l *LinkedList) Remove(data int) {
	var prev *LinkedListNode = nil
	var n *LinkedListNode = l.head
	for n != nil && n.data != data {
		prev = n
		n = n.next
	}

	if n == nil {
		return
	}

	if n == l.head {
		l.head = n.next
	} else {
		prev.next = n.next
	}

	l.count--
}

// AsArray return nodes values as array
func (l LinkedList) AsArray() []int {
	v := make([]int, l.count)
	current := l.head
	i := 0
	for current != nil {
		v[i] = current.data
		i++
		current = current.next
	}

	return v
}