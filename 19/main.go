package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
// 懂了，让两个指针保持n的间距
// 当快指针走到结尾，此时慢指针刚到倒数第n个位置
// fast =3
// 3->4->5
// 1->2->3
// [1],1
// [em 1]
// fast=1
// slow=em
//
// 1-> nil
// [1 2] 2
// 1-2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy
	// fast 先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 同速前进
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// fast走到结尾
	// 此时slow刚到倒数n
	slow.Next = slow.Next.Next
	return dummy.Next
}
