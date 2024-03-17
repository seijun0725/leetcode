func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    newList := &ListNode{}
    n := newList
    for list1 != nil || list2 != nil {
        if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
            n.Next = list1
            list1 = list1.Next
        } else {
            n.Next = list2
            list2 = list2.Next
        }
        n = n.Next
    }
    return newList.Next
}
