package main

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l3 := &ListNode{}
	curr := l3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return l3.Next
}

/*func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	result := mergeTwoLists(list1, list2)

	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
*/
