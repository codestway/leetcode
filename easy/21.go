package easy

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newH := &ListNode{}

	tmp := newH
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			tmp.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		}
		tmp = tmp.Next
	}
	for list1 != nil {
		tmp.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		tmp = tmp.Next
	}
	for list2 != nil {
		tmp.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		tmp = tmp.Next
	}

	return newH.Next
}
