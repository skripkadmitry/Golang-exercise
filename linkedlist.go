package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	res := ""
	m := &l
	for ; m != nil; m = m.Next {
		res = strconv.Itoa(m.Val) + res
	}
	return res + "\n"
}

func ItoList(n int) *ListNode {
	res := new(ListNode)
	res.Val = n % 10
	r1 := res
	r2 := new(ListNode)
	n /= 10
	for ; n > 0; n /= 10 {
		r2.Val = n % 10
		r1.Next = r2
		r1, r2 = r2, new(ListNode)
	}
	return res
}

const (
	P = 105950595250
	M = 9999999999999
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	res.Val = (l1.Val + l2.Val) % 10
	var (
		ext = (l1.Val + l2.Val) / 10
		r1  = res
		r2  = new(ListNode)
	)
	l1 = l1.Next
	l2 = l2.Next
	for {
		if ext != 0 {
			r1.Next = r2
			r2.Val = 1
		}
		if l1 == nil && l2 == nil {
			break
		}
		r1.Next = r2
		if l1 != nil {
			r2.Val += l1.Val
		}
		if l2 != nil {
			r2.Val += l2.Val
		}
		ext, r2.Val = r2.Val/10, r2.Val%10
		r1, r2 = r2, new(ListNode)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return res
}
