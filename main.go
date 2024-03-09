package main

import "fmt"

// Định nghĩa cấu trúc ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0}
	carry := 0

	for node := dummyNode; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode(carry % 10)
		carry /= 10
	}

	return dummyNode.Next
}

// Hàm main
func main() {
	// Test cases
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	sum := addTwoNumbers(l1, l2)
	for sum != nil {
		fmt.Print(sum.Val, " ")
		sum = sum.Next
	}
	fmt.Println()
}
