package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
	}
	head2 := &ListNode{
		Val: 2,
	}
	head0 := &ListNode{
		Val: 0,
	}
	headT := &ListNode{
		Val:  -4,
		Next: head2,
	}
	head.Next = head2
	head2.Next = head0
	head0.Next = headT
	fmt.Println(hasCycle(head))

}

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
// 3 2 0 -4 尾部指向 1 下标
// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	f := head
	s := head
	for s != nil && s.Next != nil {
		f = f.Next
		s = s.Next.Next
		if f == s {
			return true
		}
	}
	return false
}
