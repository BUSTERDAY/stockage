package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
	} else {
		elt := l1.Head
		for elt.Next != nil {
			elt = elt.Next
		}
		elt.Next = l2.Head
	}
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		elt := l.Head
		for elt.Next != nil {
			elt = elt.Next
		}
		elt.Next = n
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "a")
	ListPushBack(link, "b")
	ListPushBack(link, "c")
	ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	ListPushBack(link2, "e")
	ListPushBack(link2, "f")
	ListPushBack(link2, "g")
	ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	ListMerge(link, link2)
	PrintList(link)
}
