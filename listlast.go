package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil { // empty list
		return nil
	}
	newHead := l.Head
	for newHead.Next != nil { // traverse to the last node
		newHead = newHead.Next
	}
	return newHead.Data
}