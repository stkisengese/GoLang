package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}
	tail := l1.Tail     // tail for l1
	tail.Next = l2.Head // append l2 to the end of l1
	l1.Tail = l2.Tail
	// clean l2
	l2.Head = nil
	l2.Tail = nil
}
