package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2

}

func main() {
	list1 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}
	list2 := &ListNode{Val: 0, Next: nil}
	result := mergeTwoLists(list1, list2)
	fmt.Println(result.Val)

}
