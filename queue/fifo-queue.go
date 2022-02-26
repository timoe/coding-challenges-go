package queue

// Returns a new instance of a Queue with FIFO semantic base on a circular list
func NewFifoQueue() Queue {
	return &fifoQueue{l: &circularList{}}
}

type fifoQueue struct {
	l *circularList
}

func (q *fifoQueue) Enqueue(value interface{}) {
	q.l.add(value)
}

func (q *fifoQueue) Dequeue() (interface{}, error) {
	if q.l.head == nil {
		return nil, ErrQueueUnderflow
	}

	value := q.l.head.value
	if q.l.head == q.l.tail {
		q.l.head = nil
		q.l.tail = nil
	} else {
		q.l.head = q.l.head.next
	}

	return value, nil
}

func (q *fifoQueue) Front() (interface{}, error) {
	if q.l.head == nil {
		return nil, ErrQueueUnderflow
	}

	return q.l.head.value, nil
}

func (q *fifoQueue) Rear() (interface{}, error) {
	if q.l.tail == nil {
		return nil, ErrQueueUnderflow
	}

	return q.l.tail.value, nil
}
