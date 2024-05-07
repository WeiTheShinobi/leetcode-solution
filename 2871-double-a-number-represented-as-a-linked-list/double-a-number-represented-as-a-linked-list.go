func doubleIt(head *ListNode) *ListNode {
	var f func(node *ListNode) int
	f = func(node *ListNode) int {
		if node == nil {
			return 0
		}
		
		node.Val *= 2
    node.Val += f(node.Next)
		n := node.Val / 10
		node.Val %= 10
		return n
	}
	newHead := &ListNode{Next: head}
	f(newHead)
	if newHead.Val == 0 {
		return newHead.Next
	}
	return newHead
}