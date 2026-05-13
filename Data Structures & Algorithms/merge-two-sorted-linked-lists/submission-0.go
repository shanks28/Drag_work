func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy

    l1 := list1
    l2 := list2

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }

        cur = cur.Next
    }

    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }

    return dummy.Next
}