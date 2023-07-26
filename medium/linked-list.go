package medium

import "fmt"

type Node struct {
	data uint16
	next *Node
}

type LinkedList struct {
	head *Node
}

func (LL *LinkedList) appendAllVal() {
	nums := []uint16{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		newNode := &Node{data: nums[i]}
		if LL.head == nil {
			LL.head = newNode
		} else {
			current := LL.head
			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}
	}
}

func (LL *LinkedList) displayRev(curr *Node) {
	if curr == nil {
		return
	}
	LL.displayRev(curr.next)

	// value from returning call stack
	fmt.Printf("%v ", curr.data)
}

func LinkedlistReverse() {
	ll := LinkedList{}

	ll.appendAllVal()
	ll.displayRev(ll.head)
}
