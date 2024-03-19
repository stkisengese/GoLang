package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return // where reversal is not possible
	}
	var prev, current, next *NodeL
	prev = nil
	current = l.Head
	for current != nil { // reversal process
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = l.Head
}
