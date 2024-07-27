//Add Two Numbers (leetcode)
//Solved
//Medium
//Topics
//Companies
//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//Example 2:
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//Example 3:
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//Constraints:
//
//The number of nodes in each linked list is in the range [1, 100].
//0 <= Node.val <= 9
//It is guaranteed that the list represents a number that does not have leading zeros.

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	performTest(243, 564, 807)
	performTest(123, 456, 579)
}
func performTest(l1, l2, expectedResult int) {
	list1 := intToDigitList(l1)
	list2 := intToDigitList(l2)

	fmt.Printf("Test with l1=%d, l2=%d, expectedResult=%d: ", l1, l2, expectedResult)
	if listToNumber(addTwoNumbers(list1, list2, 0)) == expectedResult {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
func listToNumber(node *ListNode) int {
	result := 0
	for node != nil {
		result = result*10 + node.Val
		node = node.Next
	}
	return result
}

// Function to check if the sum of two linked lists equals the expected result
func checkSum(l1, l2 *ListNode, expectedResult int) bool {
	// Convert l1 and l2 to integers
	num1 := listToInt(l1)
	num2 := listToInt(l2)

	// Calculate the sum
	sum := num1 + num2

	// Check if the sum matches the expected result
	return sum == expectedResult
}

func listToInt(node *ListNode) int {
	result := 0
	for node != nil {
		result = result*10 + node.Val
		node = node.Next
	}
	return result
}
func intToDigitList(input int) *ListNode {
	node := &ListNode{}
	root := node

	digits := getDigits(input)

	for _, digit := range digits {
		node.Next = &ListNode{Val: digit}
		node = node.Next
	}

	return root.Next
}

func getDigits(num int) []int {
	var digits []int
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}
	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	return digits
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode, carrie int) *ListNode {
	var v1 int
	var v2 int
	var l1Next *ListNode
	var l2Next *ListNode

	if l1 == nil {
		l1Next = nil
		v1 = 0
	} else {
		l1Next = l1.Next
		v1 = l1.Val
	}

	if l2 == nil {
		l2Next = nil
		v2 = 0
	} else {
		l2Next = l2.Next
		v2 = l2.Val
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	partialSum := v1 + v2 + carrie
	newPartialSum := partialSum
	newCarrier := 0

	if partialSum > 9 {
		newCarrier = partialSum / 10
		newPartialSum = partialSum % 10
	}

	newNode := makeListNode(newPartialSum)

	newNode.Next = addTwoNumbers(l1Next, l2Next, newCarrier)
	return newNode
}

func makeListNode(n int) *ListNode {
	return &ListNode{
		Val:  n,
		Next: nil,
	}
}

func printList(root *ListNode) {
	var accumulator bytes.Buffer
	_printList(root, &accumulator)
}
func _printList(root *ListNode, accumulator *bytes.Buffer) {
	if root == nil {
		fmt.Println(accumulator.String())
		return
	}
	accumulator.WriteString(strconv.Itoa(root.Val))
	_printList(root.Next, accumulator)
}
