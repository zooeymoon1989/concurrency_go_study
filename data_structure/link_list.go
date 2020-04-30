package main

import (
	"fmt"
)

type node struct {
	next    *node
	element interface{}
}

type list struct {
	head *node
}

// 插入元素到链表中
func (l *list) Insert(key interface{}) {
	// 初始化一个节点
	n := &node{
		element: key,
	}

	if l.head == nil {
		l.head = n
		return
	}

	findNode := l.Find(key)

	// 如果是最后一个节点
	if findNode.next == nil {
		findNode.next = n
	} else {
		n.next = findNode
		findNode.next = n
	}

}

// 查找某个元素是否在链表中
func (l *list) Find(element interface{}) *node {

	currentPoint := l.head
	for currentPoint.next != nil {
		if currentPoint.element == element {
			return currentPoint
		}

		currentPoint = currentPoint.next
	}

	return currentPoint

}

func (l *list) Remove(element interface{}) {
	panic("implement me")
}

func (l *list) FindPrevious(element interface{}) {
	panic("implement me")
}

func (l *list) FindLast(element interface{}) {
	panic("implement me")
}

// 遍历显示所有的元素
func (l *list) Display() {

	current := l.head
	if l.head != nil {
		fmt.Println(current.element)
		for current.next != nil {
			current = current.next
			fmt.Println(current.element)
		}
	}

}

func (l *list) Append(*node) {

}

type ILinkList interface {
	Append(*node)
	Insert(key interface{})
	Find(element interface{}) *node
	Remove(element interface{})
	FindPrevious(element interface{})
	FindLast(element interface{})
	Display()
}

func main() {
	l := list{&node{element: 0}}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)
	l.Insert(7)


	current := l.head
	for current.next != nil {
		n := &node{element:current.element,next:current.next}
		current.next = n
		current = n.next
		if current.next == nil {
			s := &node{element:current.element}
			current.next = s
			break
		}
	}
	l.Display()

	if l.head != nil {
		reversePrint(l.head)
	}


}

func reversePrint(n *node)  {
	if n.next != nil {
		reversePrint(n.next)
	}
	fmt.Println(n.element)
}
