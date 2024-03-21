package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root} // to store level order traversal
	level := 0

	for len(queue) > 0 {
		level++
		currentLevelSize := len(queue)

		for i := 0; i < currentLevelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}
	return level
}
