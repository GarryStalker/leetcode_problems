package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedListNode := ListNode{}

	if list1.Val <= list2.Val {
		mergedListNode.Val = list1.Val
		if list1.Next != nil {
			mergedListNode.Next = mergeTwoLists(list1.Next, list2)
		} else {
			mergedListNode.Next = nil
		}
	} else {
		mergedListNode.Val = list2.Val
		if list2.Next != nil {
			mergedListNode.Next = mergeTwoLists(list1, list2.Next)
		} else {
			mergedListNode.Next = nil
		}
	}

	if list1.Next != nil {
		mergedListNode.Next = list1.Next
	} else {
		mergedListNode.Next = list2.Next
	}
	return &mergedListNode
}
