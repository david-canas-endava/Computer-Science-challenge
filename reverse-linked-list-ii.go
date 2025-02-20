/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	L := dummy
	for i := 1; i < left; i++ {
		L = L.Next
	}

	R := L.Next
	var nodes []*ListNode
	for i := left; i <= right; i++ {
		nodes = append(nodes, R)
		R = R.Next
	}

	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
	}

	L.Next = nodes[0]
	nodes[len(nodes)-1].Next = R

	return dummy.Next
}
