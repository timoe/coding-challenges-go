package twosum

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert := a.New(t)
	nums := []int{1, 2, 3, 4}

	assert.Equal([]int{2, 3}, twoSum(nums, 7))
}
