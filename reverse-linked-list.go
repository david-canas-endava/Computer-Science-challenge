/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head
    nodes := make([]*ListNode,0)
    nodes = append(nodes,temp)
	for temp.Next != nil {
        temp = temp.Next
		nodes = append(nodes,temp)
	}

    for i,j:=0,len(nodes)-1;i<j;i,j= i+1,j-1{
        nodes[i].Val, nodes[j].Val = nodes[j].Val,nodes[i].Val
    }
	return head
}