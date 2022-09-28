package main

import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSort(l *NodeI) *NodeI {
	var list List
	for l != nil {
		ListPushFront(&list, l.Data)
		l = l.Next
	}
	for list.Head != nil {
		l = listPushBack(l, list.Head.Data.(int))
		list.Head = list.Head.Next
	}
	return l
}

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		newNode := &NodeL{Data: data}
		newNode.Next, l.Head = l.Head, newNode
	}
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}
