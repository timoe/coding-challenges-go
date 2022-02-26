package queue

import (
	"math/rand"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestHeapify(t *testing.T) {
	assert := a.New(t)
	arr := []prioritisedQueueItem{
		{prio: 1, value: rand.Int()},
		{prio: 2, value: rand.Int()},
		{prio: 3, value: rand.Int()},
		{prio: 4, value: rand.Int()},
		{prio: 5, value: rand.Int()},
		{prio: 6, value: rand.Int()}}

	heapify(arr)

	assert.Equal(6, arr[0].prio)
	assert.Equal(5, arr[1].prio)
	assert.Equal(3, arr[2].prio)
	assert.Equal(4, arr[3].prio)
	assert.Equal(2, arr[4].prio)
	assert.Equal(1, arr[5].prio)
}

func TestEmptyPriorityQueue(t *testing.T) {
	assert := a.New(t)

	q := NewPriorityQueue()

	_, err := q.Front()
	assert.Equal(ErrQueueUnderflow, err)

	_, err = q.Rear()
	assert.Equal(ErrQueueOpRearNotImplemented, err)

	_, err = q.Dequeue()
	assert.Equal(ErrQueueUnderflow, err)
}

func TestPrioriatisedQueuing(t *testing.T) {
	assert := a.New(t)

	q := NewPriorityQueue()

	q.Enqueue(1, "Hello")
	q.Enqueue(2, "World")
	q.Enqueue(3, "!")

	val, err := q.Dequeue()
	assert.Nil(err)
	assert.Equal("!", val)

	val, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal("World", val)

	val, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal("Hello", val)

	_, err = q.Dequeue()
	assert.Equal(err, ErrQueueUnderflow)
}

func TestPrioriatisedEnqueingQueuing(t *testing.T) {
	assert := a.New(t)

	q := NewPriorityQueue()

	q.Enqueue(1, "Hello")
	q.Enqueue(2, "World")

	val, err := q.Dequeue()
	assert.Nil(err)
	assert.Equal("World", val)

	q.Enqueue(3, "!")

	val, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal("!", val)

	val, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal("Hello", val)

	_, err = q.Dequeue()
	assert.Equal(ErrQueueUnderflow, err)
}
