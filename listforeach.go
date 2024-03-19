package piscine

func ListForEach(l *List, f func(*NodeL)) {
	newHead := l.Head
	for newHead != nil { // iterate starting from the head
		f(newHead)             // call the function with the newHead node
		newHead = newHead.Next // move to next mode
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
