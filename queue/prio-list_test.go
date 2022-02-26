package queue

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestCircularList(t *testing.T) {
	testList(NewCircularList(), a.New(t))
}

func testList(list list, assert *a.Assertions) {
	assert.Equal(-1, list.index("42"))
	testvalues := []string{"Hello", "World", "!"}

	for _, v := range testvalues {
		list.add(v)
	}

	for i, v := range testvalues {
		index := list.index(v)
		assert.Equal(index, i, "Expecting index %d for value %s but got %d", i, v, index)
	}
}
