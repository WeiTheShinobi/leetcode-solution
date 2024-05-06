func removeNodes(head *ListNode) *ListNode {
    stack := make([]*ListNode, 0)

    for head != nil {
        for len(stack) != 0 {
            top := stack[len(stack) - 1]
            if top.Val < head.Val {
                stack = stack[:len(stack) - 1]
            } else {
                break
            }
        }
        stack = append(stack, head)
        head = head.Next
    }

    for i:=0; i<len(stack) - 1; i++ {
        stack[i].Next = stack[i + 1]
    }

    return stack[0]
}