/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var length int
	var curNode *ListNode=head
	for curNode!=nil {
		length++
		curNode=curNode.Next
	}
	var toStop int=length-n
	// which means head is being removed
	if toStop==0 {
		return head.Next
	}
	curNode=head
	var prev *ListNode=nil
	var curIndex=0
	for curIndex!=toStop {
		prev=curNode
		curNode=curNode.Next
		curIndex++
	}
	prev.Next=curNode.Next
	fmt.Println("Length is :",length)
	return head
}
