/*
Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
that node.
EXAMPLE
lnput:the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a ->b->d->e->f
pg94
*/

package main

import (
	"fmt"
)

func main() {
	l := buildListFromSlice([]int{})
	printList(l)
	printList(removeMiddleNode2(l))

	//printList(buildListFromSlice([]int{}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeMiddleNode2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := &ListNode{0, head}
	slow := prev
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return prev.Next
}

//	->
//
// 1       ->
// 1 2     -> 2
// 1 2 3   -> 1 3
// 1 2 3 4 -> 1 3 4 (4/2 = 2). Se extrae el t/2
// 1 2 3 4 5 -> 1 2 4 5 (5/2 = 2). Se extrae el t/2 + 1
func removeMiddleNode(head *ListNode) *ListNode {
	runner := head
	counter := 0

	for runner != nil {
		counter++
		runner = runner.Next
	}

	middle := counter/2 + 1

	runner = head
	prev := &ListNode{Next: head}
	counter = 0

	for runner != nil {
		counter++
		if counter == middle {
			if middle == 1 {
				head = runner.Next
			} else {
				prev.Next = runner.Next
			}
			return head
		}
		prev = prev.Next
		runner = runner.Next
	}

	return head
}

func buildListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	root := &ListNode{}
	current := root

	for n := 0; n < len(nums); n++ {
		current.Val = nums[n]
		if n < len(nums)-1 {
			current.Next = &ListNode{}
		}
		current = current.Next
	}
	return root
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d - ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
