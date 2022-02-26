package queue

import "errors"

var ErrQueueUnderflow = errors.New("queue underflow error")
var ErrQueueOpRearNotImplemented = errors.New("queue operation rear is not implemented")

type baseQueue interface {
	// Returns the value the first element of this queue and removes the element.
	// What the first element is depends on semantics of this queue, either FIFO or LIFO.
	// If queue is empty, collection.ErrQueueUnderflow is returned, returned value will be nil.
	Dequeue() (interface{}, error)

	// Returns the value of the first element from this queue.
	// What the first element is depends on semantics of this queue, either FIFO or LIFO.
	// If queue is empty, collection.ErrQueueUnderflow is returned, returned value will be nil.
	Front() (interface{}, error)

	// Returns the value from the last element from this queue.
	// What the last element is depends on semantics of this queue, either FIFO or LIFO.
	// If queue is empty, collection.ErrQueueUnderflow is returned, returned value will be nil.
	Rear() (interface{}, error)
}

type Queue interface {
	baseQueue
	// Adds a value to this Queue
	Enqueue(interface{})
}

type PriorityQueue interface {
	baseQueue
	// Adds a value to this Queue
	Enqueue(int, interface{})
}
