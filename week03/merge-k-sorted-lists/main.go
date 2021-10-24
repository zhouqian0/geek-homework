package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度 O(n*k*logk), n 为每个链表的元素个数，k 为链表的数量。
// 空间复杂度 O(logk), k 为链表数量，因为链表之间是两两合并，所以递归栈的深度为 logk。
func mergeKLists(lists []*ListNode) *ListNode {
	// 防止传入空链表数组时无法终止递归
	if len(lists) == 0 {
		return nil
	}
	// 当只剩一个元素时，返回唯一的元素
	if len(lists) == 1 {
		return lists[0]
	}
	// 分治，将链表数组对半切开
	k := len(lists) / 2
	// 合并两个切开的数组，因为传入的是递归函数，所以会一直拆分
	// 直到拆分到只剩下两个时，开始往回合并链表，最终合成一个
	return mergeList(mergeKLists(lists[:k]), mergeKLists(lists[k:]))
}

// mergeList 合并两个升序链表链表。
func mergeList(l1, l2 *ListNode) *ListNode {
	// 定义一个哨兵，用做最后的返回值
	dummy := &ListNode{}
	// 定义一个指针，之后的合并都是操作该指针
	// 如果直接操作 dummy，会导致 dummy 直接指向链表最后一个元素
	pre := dummy
	// 在两个链表都不为空的时候对比各自的头元素，将小的那个放入 pre 的 Next 指针中
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 更新 pre 的 Next 指针
			pre.Next = l1
			// 更新 l1，使其指向下一个节点
			l1 = l1.Next
		} else {
			// 更新 pre 的 Next 指针
			pre.Next = l2
			// 更新 l2，使其指向下一个节点
			l2 = l2.Next
		}
		// 更新 pre 指针，因为当前的 pre 的 Next 指针已经有了指向元素，所以需要使 pre 元素往后移动
		pre = pre.Next
	}
	// 如果 l1 还有剩下的，就将剩下的拼接到 pre 的 Next 上
	if l1 != nil {
		pre.Next = l1
	}
	// 如果 l2 还有剩下的，就将剩下的拼接到 pre 的 Next 上
	if l2 != nil {
		pre.Next = l2
	}
	// 返回哨兵节点的 Next 指针
	return dummy.Next
}
