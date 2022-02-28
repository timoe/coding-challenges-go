package addtwonums

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestTwoSumEqualSizedInputs(t *testing.T) {
	assert := a.New(t)

	one := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	two := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	result := addTwoNumbers(one, two)
	assert.NotNil(result)
	assert.Equal(7, result.Val)

	assert.NotNil(result.Next)
	assert.NotNil(0, result.Next.Val)

	assert.NotNil(result.Next.Next)
	assert.NotNil(8, result.Next.Next.Val)

	assert.Nil(result.Next.Next.Next)
}

func TestTwoSumNonEqualSizedInputs(t *testing.T) {
	assert := a.New(t)

	one := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}
	two := &ListNode{Val: 9, Next: nil}
	result := addTwoNumbers(one, two)
	assert.NotNil(result)
	assert.Equal(8, result.Val)

	assert.NotNil(result.Next)
	assert.NotNil(0, result.Next.Val)

	assert.NotNil(result.Next.Next)
	assert.NotNil(1, result.Next.Next.Val)
}
