package addtwonums

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	one := l1
	two := l2
	var carry int = 0
	var sum int = 0
	var result *ListNode = nil
	var next *ListNode = nil
	for {
		sum, carry = remainder(one.Val + two.Val + carry)

		if result == nil {
			result = &ListNode{Val: sum, Next: nil}
			next = result
		} else {
			next.Next = &ListNode{Val: sum, Next: nil}
			next = next.Next
		}

		if noneNext(one, two) {
			if carry > 0 {
				next.Next = &ListNode{Val: carry, Next: nil}
			}
			return result
		}

		one = safeNext(one)
		two = safeNext(two)
	}
}

func remainder(sum int) (int, int) {
	if sum-10 < 0 {
		return sum, 0
	}

	return sum - 10, 1
}

func noneNext(one, two *ListNode) bool {
	return one.Next == nil && two.Next == nil
}

func safeNext(node *ListNode) *ListNode {
	if node.Next == nil {
		return &ListNode{Val: 0, Next: nil}
	}

	return node.Next
}
