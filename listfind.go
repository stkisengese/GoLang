package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	newHead := l.Head    // start from head list
	for newHead != nil { //iterate thro the list
		if comp(newHead.Data, ref) { // compare data in the current mode
			return &newHead.Data // pointer to the data in the newHead mode
		} //returns when data matches
		newHead = newHead.Next // move to the next node
	}
	return nil
}
