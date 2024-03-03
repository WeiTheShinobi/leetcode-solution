/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    r := dummy
    for i := 0; i < n; i++ {
        r = r.Next
    }

    l := dummy
    for r.Next != nil {
        l = l.Next
        r = r.Next
    }

    l.Next = l.Next.Next

    return dummy.Next
}