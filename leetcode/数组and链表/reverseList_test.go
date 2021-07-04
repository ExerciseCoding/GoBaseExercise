package 数组and链表

import (
	"fmt"
	"testing"
)

/**
Leetcode: 206题：反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]
*/


type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList(head *ListNode)(*ListNode){
	if head == nil || head.Next == nil{
		return head
	}
	p,q := head.Next,head.Next.Next
	head.Next = nil
	for p != nil{
		p.Next,head,p = head,p,q
		if q != nil{
			q = q.Next
		}
	}
	return head

}

func TestReverseList(t *testing.T){
	Node3 := ListNode{
		Val:  3,
		Next: nil,
	}
	Node2 := ListNode{
		Val:  2,
		Next: &Node3,
	}

	Node1 := ListNode{
		Val:  1,
		Next: &Node2,
	}

	hea,head := &Node1,&Node1
	fmt.Print("反转前:")
	for head != nil{
		fmt.Print(head.Val," ")
		head = head.Next
	}
	fmt.Println("反转后")
	h := reverseList(hea)
	for h != nil{
		fmt.Print(h.Val," ")
		h = h.Next
	}
}