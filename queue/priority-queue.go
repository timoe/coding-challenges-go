package queue

// Returns a new instance of a PriorityQueue
func NewPriorityQueue() PriorityQueue {
	return &prioQueue{items: make([]prioritisedQueueItem, 0)}
}

type prioritisedQueueItem struct {
	prio  int
	value interface{}
}

type prioQueue struct {
	items []prioritisedQueueItem
}

func (q *prioQueue) Enqueue(prio int, value interface{}) {
	item := prioritisedQueueItem{prio: prio, value: value}
	q.items = append(q.items, item)
	heapify(q.items)
}

func (q *prioQueue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, ErrQueueUnderflow
	}
	value := q.items[0].value
	q.items = q.items[1:]
	heapify(q.items)

	return value, nil
}

func (q *prioQueue) Front() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, ErrQueueUnderflow
	}
	return q.items[0].value, nil
}

func (q *prioQueue) Rear() (interface{}, error) {
	return nil, ErrQueueOpRearNotImplemented
}

func heapify(items []prioritisedQueueItem) {
	n := len(items)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(items, n, i)
	}
}

func maxHeapify(items []prioritisedQueueItem, n, startAt int) {
	current := startAt
	left := 2*startAt + 1
	right := 2*startAt + 2

	if left < n && items[left].prio > items[current].prio {
		current = left
	}

	if right < n && items[right].prio > items[current].prio {
		current = right
	}

	if current != startAt {
		swap := items[startAt]
		items[startAt] = items[current]
		items[current] = swap

		maxHeapify(items, n, current)
	}
}
