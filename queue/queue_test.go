package queue

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestEmptyFifoQueue(t *testing.T) {
	assert := a.New(t)

	q := NewFifoQueue()

	_, err := q.Front()
	assert.Equal(err, ErrQueueUnderflow)

	_, err = q.Rear()
	assert.Equal(err, ErrQueueUnderflow)

	_, err = q.Dequeue()
	assert.Equal(err, ErrQueueUnderflow)
}

func TestBasicFifoQueueOperations(t *testing.T) {
	assert := a.New(t)

	q := NewFifoQueue()

	element1 := "Hello"
	q.Enqueue(element1)
	value, err := q.Front()
	assert.Nil(err)
	assert.Equal(value, element1)

	value, err = q.Rear()
	assert.Nil(err)
	assert.Equal(value, element1)

	value, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal(value, element1)

	_, err = q.Front()
	assert.Equal(err, ErrQueueUnderflow)

	_, err = q.Rear()
	assert.Equal(err, ErrQueueUnderflow)

	_, err = q.Dequeue()
	assert.Equal(err, ErrQueueUnderflow)
}

func TestEnqueuingDequeuingOnFifoQueue(t *testing.T) {
	assert := a.New(t)

	q := NewFifoQueue()

	element1 := "Hello"
	element2 := "World"
	element3 := "!"

	q.Enqueue(element1)
	q.Enqueue(element2)

	value, err := q.Front()
	assert.Nil(err)
	assert.Equal(element1, value)

	value, err = q.Rear()
	assert.Nil(err)
	assert.Equal(element2, value)

	value, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal(element1, value)

	q.Enqueue(element3)

	value, err = q.Rear()
	assert.Nil(err)
	assert.Equal(element3, value)

	value, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal(element2, value)

	value, err = q.Dequeue()
	assert.Nil(err)
	assert.Equal(element3, value)

	_, err = q.Front()
	assert.Equal(ErrQueueUnderflow, err)

	_, err = q.Rear()
	assert.Equal(ErrQueueUnderflow, err)

	_, err = q.Dequeue()
	assert.Equal(ErrQueueUnderflow, err)
}
