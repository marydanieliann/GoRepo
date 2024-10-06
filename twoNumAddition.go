package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	curr, carry := l3, 0

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: carry % 10}
		carry /= 10
		curr = curr.Next
	}

	return l3.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

/*func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(l1, l2)
	fmt.Print("Result of (2 -> 4 -> 3) + (5 -> 6 -> 4): ")
	printLinkedList(result)

	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}

	result = addTwoNumbers(l1, l2)
	fmt.Print("Result of (0) + (0): ")
	printLinkedList(result)

	l1 = &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}

	l2 = &ListNode{Val: 1}

	result = addTwoNumbers(l1, l2)
	fmt.Print("Result of (9 -> 9 -> 9) + (1): ")
	printLinkedList(result)

	l1 = &ListNode{Val: 1}

	l2 = &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}

	result = addTwoNumbers(l1, l2)
	fmt.Print("Result of (1) + (9 -> 9): ")
	printLinkedList(result)
}*/
