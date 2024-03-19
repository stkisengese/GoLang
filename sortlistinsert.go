package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}
	if l == nil {
		return newNode
	}
	current := l          // Find the insertion position
	prev := (*NodeI)(nil) // Cast nil to avoid type mismatch
	for current != nil && current.Data <= data_ref {
		prev = current
		current = current.Next
	}
	if prev == nil {
		newNode.Next = l // Insert at the beginning
		return newNode
	}
	prev.Next = newNode // Insert in the middle or end
	newNode.Next = current
	return l
}
