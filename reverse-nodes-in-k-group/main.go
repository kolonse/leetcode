package main

import (
	"fmt"
)

// ListNode xxxx
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil {
		return head
	}

	count := 0
	var resHead, resTail, head1, tail1, temp *ListNode
	for h := head; h != nil; {
		temp = h.Next
		if head1 == nil {
			head1 = h
			tail1 = h
		} else {
			if count < k {
				h.Next = head1
				head1 = h
			}
		}

		h = temp
		count++
		if count >= k {
			if resHead == nil {
				resHead = head1
			} else {
				resTail.Next = head1
			}

			resTail = tail1
			head1 = nil
			tail1 = nil
			count = 0
		}
	}

	if resTail == nil {
		resTail = tail1
	}
	if count > 0 {
		tail1.Next = nil
		head = head1
		head1 = nil
		tail1 = nil

		for h := head; h != nil; {
			temp = h.Next
			if head1 == nil {
				head1 = h
				tail1 = h
			} else {
				h.Next = head1
				head1 = h
			}

			h = temp
		}
		tail1.Next = nil
		if resHead == nil {
			return head1
		}
		resTail.Next = head1
		resTail = tail1
	}

	resTail.Next = nil
	return resHead
}

func main() {
	{
		node1 := &ListNode{1, nil}
		node2 := &ListNode{2, nil}
		node3 := &ListNode{3, nil}
		node4 := &ListNode{4, nil}
		node5 := &ListNode{5, nil}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		result := reverseKGroup(node1, 2)
		for h := result; h != nil; h = h.Next {
			fmt.Print(h.Val, " ")
		}
		fmt.Println()
	}

	{
		node1 := &ListNode{1, nil}
		node2 := &ListNode{2, nil}
		node3 := &ListNode{3, nil}
		node4 := &ListNode{4, nil}
		node5 := &ListNode{5, nil}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		result := reverseKGroup(node1, 3)
		for h := result; h != nil; h = h.Next {
			fmt.Print(h.Val, " ")
		}
		fmt.Println()
	}

	{
		node1 := &ListNode{1, nil}
		node2 := &ListNode{2, nil}
		node3 := &ListNode{3, nil}
		node4 := &ListNode{4, nil}
		node5 := &ListNode{5, nil}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		result := reverseKGroup(node1, -1)
		for h := result; h != nil; h = h.Next {
			fmt.Print(h.Val, " ")
		}
		fmt.Println()
	}

	{
		node1 := &ListNode{1, nil}
		node2 := &ListNode{2, nil}
		node3 := &ListNode{3, nil}
		node4 := &ListNode{4, nil}
		node5 := &ListNode{5, nil}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		result := reverseKGroup(node1, 5)
		for h := result; h != nil; h = h.Next {
			fmt.Print(h.Val, " ")
		}
		fmt.Println()
	}

	{
		node1 := &ListNode{1, nil}
		node2 := &ListNode{2, nil}
		node3 := &ListNode{3, nil}
		node4 := &ListNode{4, nil}
		node5 := &ListNode{5, nil}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		result := reverseKGroup(node1, 6)
		for h := result; h != nil; h = h.Next {
			fmt.Print(h.Val, " ")
		}
		fmt.Println()
	}
}
