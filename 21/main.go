package main

func main() {
	list11 := &ListNode{
		Val: 1,
	}
	list12 := &ListNode{
		Val: 2,
	}
	list14 := &ListNode{
		Val: 4,
	}
	list21 := &ListNode{
		Val: 1,
	}
	list23 := &ListNode{
		Val: 3,
	}
	list24 := &ListNode{
		Val: 4,
	}
	list1 := list11
	list11.Next = list12
	list12.Next = list14
	list2 := list21
	list21.Next = list23
	list23.Next = list24
	mergeTwoLists(list1, list2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 假节点，用于处理最开始不知道是哪个
	dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
			tail = tail.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
			tail = tail.Next
		}
	}
	if list1 == nil {
		tail.Next = list2
	}
	if list2 == nil {
		tail.Next = list1
	}
	return dummy.Next
	//
	//if list1 == nil {
	//	return list2
	//}
	//if list2 == nil {
	//	return list1
	//}
	//var res *ListNode
	//if list1.Val <= list2.Val {
	//	res = list1
	//	list1 = list1.Next
	//} else {
	//	res = list2
	//	list2 = list2.Next
	//}
	//var index *ListNode
	//index = res
	//for {
	//	if list1 == nil && list2 == nil {
	//		break
	//	}
	//	// 找最小节点
	//	if list1 == nil {
	//		index.Next = list2
	//		list2 = list2.Next
	//		index = index.Next
	//	} else if list2 == nil {
	//		index.Next = list1
	//		list1 = list1.Next
	//		index = index.Next
	//	} else if list1.Val <= list2.Val {
	//		index.Next = list1
	//		list1 = list1.Next
	//		index = index.Next
	//	} else {
	//		index.Next = list2
	//		list2 = list2.Next
	//		index = index.Next
	//	}
	//}
	//return res
}
