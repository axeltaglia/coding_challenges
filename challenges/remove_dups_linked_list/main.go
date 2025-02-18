/*
Remove Dups: Write code to remove duplicates from an unsorted linked list.
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
*/

package main

import "fmt"

func main() {
	l1 := giveMeAListFromArray([]int{1, 2, 3, 2, 1, 6})
	removeDups2(l1)
	printList(l1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDups2(head *ListNode) {
	current := head
	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Val == current.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}

func removeDups(head *ListNode) {
	if head == nil {
		return
	}
	hash := make(map[int]bool)
	hash[head.Val] = true
	current := head
	for current.Next != nil {
		_, ok := hash[current.Next.Val]
		if ok {
			current.Next = current.Next.Next
		} else {
			hash[current.Next.Val] = true
			current = current.Next
		}
	}
}

func giveMeAListFromArray(nums []int) *ListNode {
	head := &ListNode{}
	c := head

	for _, n := range nums {
		c.Next = &ListNode{
			Val:  n,
			Next: nil,
		}
		c = c.Next
	}

	return head.Next
}

func printList(l *ListNode) {
	if l == nil {
		fmt.Printf("\n")
		return
	}
	fmt.Printf("%d - ", l.Val)
	printList(l.Next)
}
