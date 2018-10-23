/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    resultList := &ListNode{}
    p, q, cur := l1, l2, resultList
    
    flag := 0 
    
    for { 
        if p == nil && q == nil {
            break
        }
        
        x := 0
        y := 0
        
        if p != nil {x = p.Val}
        if q != nil {y = q.Val}
        
        sum := x + y + flag
        flag = sum / 10
        single := sum % 10
        
        cur.Next = &ListNode{Val: single}
        cur = cur.Next
        
        if p != nil {
            p = p.Next
        }
        
        if q != nil {
            q = q.Next
        }
    } 
    
    if flag == 1 {   //only l1,l2 have same length and sum of the last two num large than 10
        cur.Next = &ListNode{Val: 1}
    }
    
    return resultList.Next
}
