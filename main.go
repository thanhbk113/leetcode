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
func reorderList(head *ListNode) {
	if head == nil && head.Next == nil {
		return
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := reverseList(slow.Next)
	slow.Next = nil

	p1 := head
	p2 := secondHalf

	for p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p1.Next = p2
		p2.Next = next1
		p1 = next1
		p2 = next2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// Hàm main
func main() {
	head := &ListNode{Val: 1}
	current := head
	for i := 2; i <= 5; i++ {
		node := &ListNode{Val: i}
		current.Next = node
		current = node
	}

	fmt.Println("Danh sách liên kết ban đầu:")
	printList(head)

	reorderList(head)

	fmt.Println("Danh sách liên kết sau khi sắp xếp lại:")
	printList(head)
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
