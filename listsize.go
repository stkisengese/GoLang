package piscine
/*
type NodeL struct {
  Data interface{}
  Next *NodeL
}

type List struct {
  Head *NodeL
  Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
} */

func ListSize(l *List) int {
  	size := 0
  	newHead := l.Head
  	for newHead != nil {
    	   size++
       	newHead = newHead.Next 
  	}
  	return size
}
