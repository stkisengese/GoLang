package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	newNode := l.Head
	for newNode != nil {
		if newNode.Data == data_ref {
			if prev == nil {
				l.Head = newNode.Next
			} else {
				prev.Next = newNode.Next
				if newNode == l.Tail {
					l.Tail = prev
				}
			}
			newNode = newNode.Next
		} else {
			prev = newNode
			newNode = newNode.Next
		}
	}
}
