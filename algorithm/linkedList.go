package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func (ll *LinkedList) Insert(n int) {
	node := &Node{Data: n}
	if ll.Head == nil {
		ll.Head = node
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (ll *LinkedList) Delete(n int) {
	if ll.Head == nil {
		fmt.Println("Empty LL.")
		return
	}

	if ll.Head.Data == n {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == n {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (ll *LinkedList) Contain(n int) bool {
	if ll.Head.Data == n {
		ll.Head = ll.Head.Next
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == n {
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	ll.Display()
	fmt.Println("ll.Contain(3)", ll.Contain(3))
	fmt.Println("ll.Contain(5)", ll.Contain(5))

	ll.Delete(1)
	fmt.Println("ll.Contain(1)", ll.Contain(1))
	ll.Display()
}
