/**
* Swap Nodes in Pairs - https://leetcode.com/problems/swap-nodes-in-pairs/
*/

package main

import (
	"leetcode/def"
)

func swapPairsUnknow(head *def.ListNode) *def.ListNode {
    var (
    	pre = &def.ListNode{0 , nil}
    	res = pre
    )
    for head != nil && head.Next != nil {
    	next := head.Next.Next
    	pre.Next = head.Next
    	pre = head
    	head.Next.Next = head
    	head = next
    }
    pre.Next = head
    return res.Next
}

func reverseKGroup(head *def.ListNode, k int) *def.ListNode {
    var (
    	temp = head
    	n = 0
    	pre = &def.ListNode{0 , nil}
    )
    for temp != nil {
    	n ++
    	temp = temp.Next
    }
    temp = pre
    for head != nil {
    	if n >= k {
    		next , cur := reverseNPairs(head, k)
    		temp.Next = next
	    	for temp.Next != nil {
	    		temp = temp.Next
	    	}
	    	head = cur
	    	n -= k
    	}else {
    		temp.Next = head
    		break
    	}
    }
    return pre.Next
}

func reverseNPairs(head *def.ListNode , n int) (*def.ListNode , *def.ListNode){
	if head == nil || head.Next == nil || n < 1 {
		return head , nil
	}
	var (
		pre = head
		cur = head.Next
		i = 1
	)
	for cur != nil {
		if i == n {
			break
		}
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		i ++
	}
	head.Next = nil
	return pre , cur
}

func swapNPairs(head *def.ListNode , n int) *def.ListNode{
	var (
    	pre = &def.ListNode{0 , nil}
    	temp = pre
    )
    for head != nil {
    	next , cur := reverseNPairs(head, n)
    	temp.Next = next
    	for temp.Next != nil {
    		temp = temp.Next
    	}
    	head = cur
    }
	return pre.Next
}

func swapPairs(head *def.ListNode) *def.ListNode {
    var (
    	pre = &def.ListNode{0 , nil}
    	temp = pre
    )
    pre.Next = head
    for temp.Next != nil && temp.Next.Next != nil {
    	p := temp.Next
    	q := temp.Next.Next
    	temp.Next = q
    	p.Next = q.Next
    	q.Next = p
    	temp = p
    }
    return pre.Next
}

func NewHead() *def.ListNode {
	head := def.NewNode(1)
	temp := head
	for i := 2 ; i <= 5 ; i++ {
		if temp.Next == nil {
			temp.Next = def.NewNode(i)
		}
		temp = temp.Next
	}
	return head	
}

func main() {
	head := NewHead()
	head.Show()
	// swapPairs(head).Show()
	// swapNPairs(head, 1).Show()
	reverseKGroup(head, 3).Show()
}