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
	copy := head
	temp := copy
	depth := 0
	for temp.Next != nil {
		depth++
		temp = temp.Next
	}
	r := temp
    RHead := r
	for x := depth; x >= 1; x-- {
		temp = copy
		for y := x - 1; y >= 1; y-- {
			temp = temp.Next
		}
		temp.Next = nil
		r.Next = temp
		r = r.Next
	}

	return RHead
}