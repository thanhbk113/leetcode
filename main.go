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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 != nil && list2 != nil {
		return nil
	}

	merge := &ListNode{}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			Append(merge, list1)
		} else {
			Append(merge, list2)
		}
		merge = merge.Next
		list1 = list1.Next
		list2 = list2.Next
	}
	return merge
}

func Append(merge *ListNode, node *ListNode) {
	if merge == nil {
		merge = node
		merge.Next = nil
	}
	merge.Next = node
	merge.Next.Next = nil
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	n := 4
	curr := head
	for i := 2; i <= n; i++ {
		newNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		curr.Next = newNode
		curr = newNode
	}
}
