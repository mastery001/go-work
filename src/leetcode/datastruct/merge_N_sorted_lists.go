/**
* Merge Two Sorted Lists - https://leetcode.com/problems/merge-two-sorted-lists/
*/

/**
* Merge k Sorted Lists - https://leetcode.com/problems/merge-k-sorted-lists/
* 1. brute force 
* 2. priority queue
* 3. division
*/

package main

import (
	"leetcode/def"
)

func mergeKLists(lists []*def.ListNode) *def.ListNode {
    n := len(lists)
    if n == 0 {
    	return nil
    } else if n == 1 {
		return lists[0]
	}else {
		mid := n >> 1
		return mergeTwoLists(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:n]))
	}
}

func mergeTwoLists(l1 *def.ListNode, l2 *def.ListNode) *def.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		res *def.ListNode = &def.ListNode{-1 , nil}
	)
	temp := res
    for l1 != nil && l2 != nil {
    	if l1.Val <= l2.Val {
    		temp.Next = l1
    		l1 = l1.Next
    	}else {
    		temp.Next = l2
    		l2 = l2.Next
    	}
    	temp = temp.Next
    }
    if l1 != nil {
    	temp.Next = l1
    }
    if l2 != nil {
    	temp.Next = l2
    }
    return res.Next
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
	// mergeTwoLists(NewHead() ,NewHead()).Show()
	mergeKLists([]*def.ListNode{NewHead()}).Show()
}