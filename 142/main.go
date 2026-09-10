package main

import "fmt"

func main() {
	head3 := &ListNode{
		Val: 3,
	}
	head2 := &ListNode{
		Val: 2,
	}
	head0 := &ListNode{
		Val: 0,
	}
	head4 := &ListNode{
		Val: -4,
	}
	head3.Next = head2
	head2.Next = head0
	head0.Next = head4
	head4.Next = head2
	printRes := func(head *ListNode) string {
		if head == nil {
			return "null"
		}
		return fmt.Sprintf("{%d}", head.Val)
	}
	fmt.Println(printRes(detectCycle(head3)))
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
// head = [3,2,0,-4], pos = 1
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 找到相遇的点了
			break
		}
	}
	if fast != slow {
		return nil
	}
	for head != nil && head.Next != nil {
		if head == slow {
			return head
		}
		head = head.Next
		slow = slow.Next
	}
	return nil
}
