package clonegraph

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	assert := a.New(t)
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}

	one.Neighbors = append(one.Neighbors, two, four)
	two.Neighbors = append(two.Neighbors, one, three)
	three.Neighbors = append(three.Neighbors, two, four)
	four.Neighbors = append(four.Neighbors, one, three)

	oneCloned := cloneGraph(one)

	assert.False(oneCloned == one)
	assert.Equal(1, oneCloned.Val)
	assert.Equal(2, len(oneCloned.Neighbors))

	twoCloned := oneCloned.Neighbors[0]
	assert.False(twoCloned == two)
	assert.Equal(2, twoCloned.Val)
	assert.Equal(2, len(twoCloned.Neighbors))

	fourCloned := oneCloned.Neighbors[1]
	assert.False(fourCloned == four)
	assert.Equal(4, fourCloned.Val)
	assert.Equal(2, len(fourCloned.Neighbors))

	threeCloned := twoCloned.Neighbors[1]
	assert.False(twoCloned.Neighbors[1] == three)
	assert.Equal(3, threeCloned.Val)
}

func TestCloneOneNodeGraph(t *testing.T) {
	assert := a.New(t)
	one := &Node{Val: 1}
	oneCloned := cloneGraph(one)

	assert.NotNil(oneCloned)
	assert.False(oneCloned == one)
}

func TestCloneNil(t *testing.T) {
	assert := a.New(t)
	oneCloned := cloneGraph(nil)

	assert.Nil(oneCloned)
}
