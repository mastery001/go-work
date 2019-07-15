/**
* Remove Nth Node From End of List - https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"leetcode/def"
) 
 
func removeNthFromEnd(head *def.ListNode, n int) *def.ListNode {
    var (
    	root = &def.ListNode{0 , head}
    	p = root
    	q = root
    	i = 0
    	end = n + 1
    )
    for q != nil {
    	if i >= end {
    		p = p.Next
    	}
    	q = q.Next
    	i ++
    }
 	if p.Next.Next == nil {
 		p.Next = nil
 	}else  {
 		p.Next = p.Next.Next
 	}

    return root.Next
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

	removeNthFromEnd(NewHead() , 1).Show()
	removeNthFromEnd(NewHead() , 2).Show()
	removeNthFromEnd(NewHead() , 3).Show()
	removeNthFromEnd(NewHead() , 4).Show()
	removeNthFromEnd(NewHead() , 5).Show()
}