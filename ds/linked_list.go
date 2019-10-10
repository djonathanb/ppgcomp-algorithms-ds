package ds

// LinkedListNode stores data
type LinkedListNode struct {
	key  int
	data interface{}
	next *LinkedListNode
}

// NewLinkedListNode creates a new Node storing data
func NewLinkedListNode(key int, data interface{}) LinkedListNode {
	return LinkedListNode{
		key:  key,
		data: data,
		next: nil,
	}
}

// Key returns the node key
func (n LinkedListNode) Key() int {
	return n.key
}

// Data returns de node stored data
func (n LinkedListNode) Data() interface{} {
	return n.data
}

// LinkedList list of nodes
type LinkedList struct {
	head  *LinkedListNode
	count int
}

// NewLinkedList creates a new empty Liskend List
func NewLinkedList() LinkedList {
	return LinkedList{
		head:  nil,
		count: 0,
	}
}

// Insert insert a new item at the head of structure
func (l *LinkedList) Insert(key int, data interface{}) {
	node := NewLinkedListNode(key, data)
	node.next = l.head
	l.head = &node
	l.count++
}

// Each executes func for each (index, node) in the list
func (l LinkedList) Each(fn func(int, LinkedListNode)) {
	i := 0
	node := l.head
	for node != nil {
		fn(i, *node)
		node = node.next
	}
}

// Search return a node given your key
func (l *LinkedList) Search(key int) *LinkedListNode {
	current := l.head
	for current != nil && current.key != key {
		current = current.next
	}
	return current
}

// Remove remove the given item from structure
func (l *LinkedList) Remove(key int) {
	var prev *LinkedListNode = nil
	var n *LinkedListNode = l.head
	for n != nil && n.key != key {
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

// AsArray return nodes keys as array
func (l LinkedList) AsArray() []int {
	v := make([]int, l.count)
	current := l.head
	i := 0
	for current != nil {
		v[i] = current.key
		i++
		current = current.next
	}

	return v
}
