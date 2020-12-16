package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	list := lists[0]
	length := len(lists)

	head := &ListNode{
		-10001,
		list,
	}

	for i := 1; i < length; i++ {
		father := head
		n1 := father.Next
		n2 := lists[i]

		for n1 != nil && n2 != nil {
			if n1.Val > n2.Val {
				tmp := n2.Next
				father.Next = n2
				n2.Next = n1
				father = n2
				n2 = tmp
			} else {
				father = n1
				n1 = n1.Next
			}
		}

		if n1 == nil && n2 != nil {
			father.Next = n2
		}
	}
	return head.Next
}

func main() {
	lists := []*ListNode{
		nil,
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{4, &ListNode{8, nil}}},
		&ListNode{2, &ListNode{3, &ListNode{9, nil}}},
		&ListNode{2, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
	}

	results := mergeKLists(lists)

	for node := results; node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}
	fmt.Println()
}
