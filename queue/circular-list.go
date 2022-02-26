package queue

// abstracts a list, used internally to not lead circularList struct
type list interface {
	add(interface{})
	index(interface{}) int
}

type circularListNode struct {
	value interface{}
	next  *circularListNode
	prev  *circularListNode
}

// implements a list which is circular
type circularList struct {
	head *circularListNode
	tail *circularListNode
}

// Create a circualr List. For this implementation, it maintains
// additional pointer for previous nodes.
func NewCircularList() list {
	return &circularList{}
}

func (l *circularList) add(value interface{}) {
	if l.head == nil {
		node := &circularListNode{value: value}
		l.head = node
		l.head.next = node
		l.head.prev = node
		l.tail = node
		return
	}

	node := &circularListNode{value: value}
	l.tail.next = node
	node.next = l.head
	node.prev = l.tail
	l.tail = node
}

func (l *circularList) index(value interface{}) int {
	if l.head == nil {
		return -1
	}

	index := 0
	node := l.head
	for {
		if node.value == value {
			return index
		}
		node = node.next
		index++

		if node == l.head {
			break
		}
	}

	return -1
}
