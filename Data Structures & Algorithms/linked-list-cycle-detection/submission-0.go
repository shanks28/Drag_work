/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	set:=make(map[*ListNode]int)
	var cur *ListNode=head
	if cur == nil {
		return false
	}
	for cur != nil {
		if _,ok:=set[cur]; ok{
			return true
		}
		set[cur]=cur.Val
		cur=cur.Next
	}
	return false
	
    
}
