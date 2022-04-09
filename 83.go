package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for i := head; i.Next != nil; {
		if i.Val == i.Next.Val {
			if i.Next.Next == nil {
				i.Next = nil
			} else {
				i.Next = i.Next.Next
			}
		} else {
			i = i.Next
		}
	}
	return head

}
