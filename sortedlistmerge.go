package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	temp := &NodeI{0, nil}
	tail := temp
	current1 := n1
	current2 := n2
	for current1 != nil && current2 != nil {
		if current1.Data <= current2.Data {
			tail.Next = current1
			current1 = current1.Next
		} else {
			tail.Next = current2
			current2 = current2.Next
		}
		tail = tail.Next
	}
	tail.Next = current1
	if current2 != nil {
		tail.Next = current2
	}
	return temp.Next
}
