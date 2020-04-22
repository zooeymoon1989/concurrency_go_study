package main

type node struct {
	prev    *node
	next    *node
	element interface{}
}

type list struct {
	head *node
	tail *node
}

func (l *list) Insert(key interface{}) {
	n := &node{
		next:    l.head,
		element: key,
	}

	if ok := l.Find(key); ok {

	}

}

func (l *list) Find(element interface{}) bool {

	currentPoint := l.head
	for currentPoint.next != nil {
		if currentPoint.element == element {
			return true
		}

		currentPoint = currentPoint.next
	}

	return false

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

func (l *list) Display() {
	panic("implement me")
}

type ILinkList interface {
	Insert(key interface{})
	Find(element interface{}) bool
	Remove(element interface{})
	FindPrevious(element interface{})
	FindLast(element interface{})
	Display()
}

func main() {

}
