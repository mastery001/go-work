package def

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
}

func (node *ListNode) Show() {
	temp := node
	for temp != nil {
		fmt.Print(temp.Val , "->")
		temp = temp.Next
	}
	fmt.Println()
}

func NewNode(val int) *ListNode {
	var head *ListNode = new(ListNode)
	head.Val = val
	return head
}



