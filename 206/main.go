package main

import "fmt"

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	var curr *ListNode
	curr = head
	for i := 2; i < 6; i++ {
		next := &ListNode{
			Val: i,
		}
		curr.Next = next
		curr = next
	}
	printRes := func(head *ListNode) string {
		str := ""
		for head != nil {
			str += fmt.Sprintf("{%d}", head.Val)
			head = head.Next
		}
		return str
	}
	fmt.Println(printRes(reverseList(head)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表处理
// [1 2 3 4 5]
// 特点，当head.next 的指针被移动，就丢失后续的了
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
