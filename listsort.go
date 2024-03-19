package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var head, newNode, nextNode *NodeI
	temp := 0
	head = l
	if head == nil {
		return nil
	}
	for newNode = head; newNode != nil; newNode = newNode.Next {
		for nextNode = newNode.Next; nextNode != nil; nextNode = nextNode.Next {
			if newNode.Data > nextNode.Data {
				temp = newNode.Data
				newNode.Data = nextNode.Data
				nextNode.Data = temp
			}
		}
	}
	return head
}
