package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	walker := head
	runner := head

	for runner.Next != nil && runner.Next.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	c := head
	for c != nil {
		_, ok := hash[c]
		if ok {
			return true
		}
		hash[c] = true
		c = c.Next
	}

	return false
}
