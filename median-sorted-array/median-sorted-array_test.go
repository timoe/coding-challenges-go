package med_sorted_array

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestMedSortedArrayUnequalArray(t *testing.T) {
	assert := a.New(t)
	assert.Equal(2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func TestMedSortedArrayEqualArray(t *testing.T) {
	assert := a.New(t)
	assert.Equal(2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
